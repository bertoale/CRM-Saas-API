package user

import (
	"crm/internal/tenant"
	"crm/pkg/config"
	"errors"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type Service interface {
	// tenant-scoped operations
	Create(req UserRequest, currentUser *User) (*UserResponse, error)
	GetAll(currentUser *User) ([]UserResponse, error)
	GetByID(id uint, currentUser *User) (*UserResponse, error)
	Update(id uint, req UserRequest, currentUser *User) (*UserResponse, error)
	Delete(id uint, currentUser *User) error
	// auth
	Login(req UserLoginRequest) (*UserLoginResponse, error)
	Logout(token string) error
	// tenant creation with owner
	CreateTenantWithOwner(tenantReq tenant.TenantRequest, userReq UserRequest) (*CreateTenantWithOwnerResponse, error)
}

type service struct {
	repo       Repository
	tenantRepo tenant.Repository
	cfg        *config.Config
	db         *gorm.DB
}

// CreateTenantWithOwner implements Service.
// This creates a new tenant and automatically creates an owner user for that tenant
func (s *service) CreateTenantWithOwner(tenantReq tenant.TenantRequest, userReq UserRequest) (*CreateTenantWithOwnerResponse, error) {
	// Check if domain already exists
	_, err := s.tenantRepo.FindByDomain(tenantReq.Domain)
	if err == nil {
		return nil, errors.New("domain sudah digunakan")
	}

	// Check if email already exists
	_, err = s.repo.FindByEmail(userReq.Email)
	if err == nil {
		return nil, errors.New("email sudah digunakan")
	}

	// Start transaction
	// All query after this not really saved to the database until tx.Commit() is called
	tx := s.db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	// Create tenant
	newTenant := &tenant.Tenant{
		Name:   tenantReq.Name,
		Domain: tenantReq.Domain,
		Status: tenant.StatusActive,
	}

	if err := tx.Create(newTenant).Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	// Hash password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(userReq.Password), bcrypt.DefaultCost)
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	// Create owner user
	newUser := &User{
		TenantID: newTenant.ID,
		Name:     userReq.Name,
		Email:    userReq.Email,
		Password: string(hashedPassword),
		Role:     RoleOwner,
		Status:   StatusActive,
	}

	if err := tx.Create(newUser).Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	// Commit transaction
	// If all good, save to database
	if err := tx.Commit().Error; err != nil {
		return nil, err
	}

	return &CreateTenantWithOwnerResponse{
		Tenant: *tenant.ToTenantResponse(newTenant),
		User:   *ToUserResponse(newUser),
	}, nil
}

// Create implements Service.
func (s *service) Create(req UserRequest, currentUser *User) (*UserResponse, error) {
	// Only owner and admin can create users
	if currentUser.Role != RoleOwner && currentUser.Role != RoleAdmin {
		return nil, errors.New("anda tidak memiliki akses untuk membuat user")
	}

	// Check if email already exists
	_, err := s.repo.FindByEmail(req.Email)
	if err == nil {
		return nil, errors.New("email sudah digunakan")
	}

	// Hash password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	// Set default role if not provided
	role := req.Role
	if role == "" {
		role = RoleSales
	}

	// Owner can only be created during tenant creation
	if role == RoleOwner {
		return nil, errors.New("role owner hanya bisa dibuat saat pembuatan tenant")
	}

	newUser := &User{
		TenantID: currentUser.TenantID,
		Name:     req.Name,
		Email:    req.Email,
		Password: string(hashedPassword),
		Role:     RoleENUM(role),
		Status:   StatusActive,
	}

	if err := s.repo.Create(newUser); err != nil {
		return nil, err
	}

	return ToUserResponse(newUser), nil
}

// Delete implements Service.
func (s *service) Delete(id uint, currentUser *User) error {
	// Only owner and admin can delete users
	if currentUser.Role != RoleOwner && currentUser.Role != RoleAdmin {
		return errors.New("anda tidak memiliki akses untuk menghapus user")
	}

	user, err := s.repo.FindByID(id, currentUser.TenantID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errors.New("user tidak ditemukan")
		}
		return err
	}

	// Cannot delete owner
	if user.Role == RoleOwner {
		return errors.New("owner tidak dapat dihapus")
	}

	// Cannot delete yourself
	if user.ID == currentUser.ID {
		return errors.New("anda tidak dapat menghapus akun sendiri")
	}

	return s.repo.Delete(id, currentUser.TenantID)
}

// GetAll implements Service.
func (s *service) GetAll(currentUser *User) ([]UserResponse, error) {
	users, err := s.repo.FindByTenantID(currentUser.TenantID)
	if err != nil {
		return nil, err
	}

	var responses []UserResponse
	for _, u := range users {
		responses = append(responses, *ToUserResponse(&u))
	}
	return responses, nil
}

// GetByID implements Service.
func (s *service) GetByID(id uint, currentUser *User) (*UserResponse, error) {
	user, err := s.repo.FindByID(id, currentUser.TenantID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("user tidak ditemukan")
		}
		return nil, err
	}
	return ToUserResponse(user), nil
}

// Login implements Service.
func (s *service) Login(req UserLoginRequest) (*UserLoginResponse, error) {
	user, err := s.repo.FindByEmail(req.Email)
	if err != nil {
		return nil, errors.New("email atau password salah")
	}

	// Check password
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)); err != nil {
		return nil, errors.New("email atau password salah")
	}
	// Check if user is active
	if user.Status != StatusActive {
		return nil, errors.New("akun anda tidak aktif")
	}

	// Generate JWT token
	tokenString, err := GenerateToken(user, s.cfg)
	if err != nil {
		return nil, err
	}

	return &UserLoginResponse{
		Token: tokenString,
		User:  *ToUserResponse(user),
	}, nil
}

// Logout implements Service.
func (s *service) Logout(token string) error {
	return nil
}

// Update implements Service.
func (s *service) Update(id uint, req UserRequest, currentUser *User) (*UserResponse, error) {
	user, err := s.repo.FindByID(id, currentUser.TenantID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("user tidak ditemukan")
		}
		return nil, err
	}

	// Only owner and admin can update other users
	// Users can update themselves
	if currentUser.ID != id && currentUser.Role != RoleOwner && currentUser.Role != RoleAdmin {
		return nil, errors.New("anda tidak memiliki akses untuk mengupdate user ini")
	}

	// Update fields
	user.Name = req.Name

	// Check if email is being changed and if it's already in use
	if user.Email != req.Email {
		existingUser, err := s.repo.FindByEmail(req.Email)
		if err == nil && existingUser.ID != id {
			return nil, errors.New("email sudah digunakan")
		}
		user.Email = req.Email
	}

	// Update password if provided
	if req.Password != "" {
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
		if err != nil {
			return nil, err
		}
		user.Password = string(hashedPassword)
	}

	// Only owner and admin can update roles
	if req.Role != "" && (currentUser.Role == RoleOwner || currentUser.Role == RoleAdmin) {
		// Cannot change owner role
		if user.Role == RoleOwner {
			return nil, errors.New("role owner tidak dapat diubah")
		}
		// Cannot set role to owner
		if req.Role == RoleOwner {
			return nil, errors.New("tidak dapat mengubah role menjadi owner")
		}
		user.Role = RoleENUM(req.Role)
	}

	if err := s.repo.Update(user, currentUser.TenantID); err != nil {
		return nil, err
	}
	return ToUserResponse(user), nil
}

func NewService(repo Repository, tenantRepo tenant.Repository, cfg *config.Config, db *gorm.DB) Service {
	return &service{
		repo:       repo,
		tenantRepo: tenantRepo,
		cfg:        cfg,
		db:         db,
	}
}
