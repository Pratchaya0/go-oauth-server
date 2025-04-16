-- Users table to store basic user information
CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    username VARCHAR(255) NOT NULL UNIQUE,
    email VARCHAR(255) NOT NULL UNIQUE,
    password_hash VARCHAR(255) NOT NULL,
    first_name VARCHAR(100),
    last_name VARCHAR(100),
    is_active BOOLEAN DEFAULT TRUE,
    is_verified BOOLEAN DEFAULT FALSE,
    verification_token VARCHAR(255),
    verification_token_expires_at TIMESTAMP,
    reset_password_token VARCHAR(255),
    reset_password_token_expires_at TIMESTAMP,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Roles table for user role definitions
CREATE TABLE roles (
    id SERIAL PRIMARY KEY,
    name VARCHAR(100) NOT NULL UNIQUE,
    description TEXT,
    is_system_role BOOLEAN DEFAULT FALSE,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- User-Role assignments
CREATE TABLE user_roles (
    id SERIAL PRIMARY KEY,
    user_id INTEGER REFERENCES users(id) ON DELETE CASCADE,
    role_id INTEGER REFERENCES roles(id) ON DELETE CASCADE,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    UNIQUE (user_id, role_id)
);

-- Permissions table for granular access control
CREATE TABLE permissions (
    id SERIAL PRIMARY KEY,
    name VARCHAR(100) NOT NULL UNIQUE,
    description TEXT,
    resource VARCHAR(100) NOT NULL,
    action VARCHAR(50) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    UNIQUE (resource, action)
);

-- Role-Permission assignments for customizable role capabilities
CREATE TABLE role_permissions (
    id SERIAL PRIMARY KEY,
    role_id INTEGER REFERENCES roles(id) ON DELETE CASCADE,
    permission_id INTEGER REFERENCES permissions(id) ON DELETE CASCADE,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    UNIQUE (role_id, permission_id)
);

-- Clients/Applications that can request access to the API
CREATE TABLE oauth_clients (
    id SERIAL PRIMARY KEY,
    client_id VARCHAR(100) NOT NULL UNIQUE,
    client_secret VARCHAR(255) NOT NULL,
    name VARCHAR(255) NOT NULL,
    description TEXT,
    website_url VARCHAR(255),
    redirect_uris TEXT[], -- Array of allowed redirect URIs
    allowed_grant_types TEXT[], -- Array of allowed OAuth grant types
    is_confidential BOOLEAN DEFAULT TRUE, -- Whether client can keep secrets (server-side apps)
    user_id INTEGER REFERENCES users(id) ON DELETE SET NULL, -- Owner of the client
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- OAuth access tokens
CREATE TABLE oauth_access_tokens (
    id SERIAL PRIMARY KEY,
    access_token VARCHAR(255) NOT NULL UNIQUE,
    client_id INTEGER REFERENCES oauth_clients(id) ON DELETE CASCADE,
    user_id INTEGER REFERENCES users(id) ON DELETE CASCADE,
    scopes TEXT[], -- Array of permission scopes
    expires_at TIMESTAMP NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- OAuth refresh tokens for obtaining new access tokens
CREATE TABLE oauth_refresh_tokens (
    id SERIAL PRIMARY KEY,
    refresh_token VARCHAR(255) NOT NULL UNIQUE,
    access_token_id INTEGER REFERENCES oauth_access_tokens(id) ON DELETE CASCADE,
    is_revoked BOOLEAN DEFAULT FALSE,
    expires_at TIMESTAMP NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Authorization codes (for authorization_code grant type)
CREATE TABLE oauth_authorization_codes (
    id SERIAL PRIMARY KEY,
    authorization_code VARCHAR(255) NOT NULL UNIQUE,
    client_id INTEGER REFERENCES oauth_clients(id) ON DELETE CASCADE,
    user_id INTEGER REFERENCES users(id) ON DELETE CASCADE,
    redirect_uri TEXT NOT NULL,
    scopes TEXT[],
    expires_at TIMESTAMP NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- API keys for simpler authentication methods
CREATE TABLE api_keys (
    id SERIAL PRIMARY KEY,
    user_id INTEGER REFERENCES users(id) ON DELETE CASCADE,
    api_key VARCHAR(100) NOT NULL UNIQUE, -- Public key
    api_secret VARCHAR(255) NOT NULL, -- Secret key (should be stored securely)
    name VARCHAR(255) NOT NULL, -- Name/purpose of this key
    is_active BOOLEAN DEFAULT TRUE,
    scopes TEXT[], -- Array of permission scopes
    expires_at TIMESTAMP,
    last_used_at TIMESTAMP,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Service table to track different web services in your ecosystem
CREATE TABLE services (
    id SERIAL PRIMARY KEY,
    name VARCHAR(100) NOT NULL UNIQUE,
    description TEXT,
    base_url VARCHAR(255),
    is_active BOOLEAN DEFAULT TRUE,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Service-specific permissions
CREATE TABLE service_permissions (
    id SERIAL PRIMARY KEY,
    service_id INTEGER REFERENCES services(id) ON DELETE CASCADE,
    permission_id INTEGER REFERENCES permissions(id) ON DELETE CASCADE,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    UNIQUE (service_id, permission_id)
);

-- OAuth scope definitions
CREATE TABLE oauth_scopes (
    id SERIAL PRIMARY KEY,
    name VARCHAR(100) NOT NULL UNIQUE,
    description TEXT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Map scopes to permissions
CREATE TABLE scope_permissions (
    id SERIAL PRIMARY KEY,
    scope_id INTEGER REFERENCES oauth_scopes(id) ON DELETE CASCADE,
    permission_id INTEGER REFERENCES permissions(id) ON DELETE CASCADE,
    UNIQUE (scope_id, permission_id)
);

-- Session tracking for users
CREATE TABLE sessions (
    id SERIAL PRIMARY KEY,
    user_id INTEGER REFERENCES users(id) ON DELETE CASCADE,
    token VARCHAR(255) NOT NULL UNIQUE,
    ip_address VARCHAR(45),
    user_agent TEXT,
    is_valid BOOLEAN DEFAULT TRUE,
    expires_at TIMESTAMP NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    last_activity_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Audit log for security events
CREATE TABLE audit_logs (
    id SERIAL PRIMARY KEY,
    user_id INTEGER REFERENCES users(id) ON DELETE SET NULL,
    client_id INTEGER REFERENCES oauth_clients(id) ON DELETE SET NULL,
    action VARCHAR(100) NOT NULL,
    resource_type VARCHAR(100),
    resource_id VARCHAR(100),
    ip_address VARCHAR(45),
    user_agent TEXT,
    details JSONB,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Create indexes for performance
CREATE INDEX idx_user_roles_user_id ON user_roles(user_id);
CREATE INDEX idx_user_roles_role_id ON user_roles(role_id);
CREATE INDEX idx_role_permissions_role_id ON role_permissions(role_id);
CREATE INDEX idx_access_tokens_user_id ON oauth_access_tokens(user_id);
CREATE INDEX idx_api_keys_user_id ON api_keys(user_id);
CREATE INDEX idx_audit_logs_user_id ON audit_logs(user_id);
CREATE INDEX idx_audit_logs_created_at ON audit_logs(created_at);

-- Insert default roles
INSERT INTO roles (name, description, is_system_role) VALUES 
('admin', 'System administrator with full access', TRUE),
('user', 'Standard user role', TRUE);

-- Insert some basic permissions
INSERT INTO permissions (name, description, resource, action) VALUES
('user:read', 'View user profiles', 'user', 'read'),
('user:create', 'Create new users', 'user', 'create'),
('user:update', 'Update user profiles', 'user', 'update'),
('user:delete', 'Delete users', 'user', 'delete'),
('role:read', 'View roles', 'role', 'read'),
('role:create', 'Create new roles', 'role', 'create'),
('role:update', 'Update roles', 'role', 'update'),
('role:delete', 'Delete roles', 'role', 'delete'),
('permission:read', 'View permissions', 'permission', 'read'),
('permission:assign', 'Assign permissions to roles', 'permission', 'assign');

-- Assign permissions to default roles
INSERT INTO role_permissions (role_id, permission_id)
SELECT 
    (SELECT id FROM roles WHERE name = 'admin'),
    id
FROM permissions;

INSERT INTO role_permissions (role_id, permission_id)
SELECT 
    (SELECT id FROM roles WHERE name = 'user'),
    id
FROM permissions
WHERE name IN ('user:read');

-- Create default OAuth scopes
INSERT INTO oauth_scopes (name, description) VALUES
('profile', 'Access to user profile information'),
('email', 'Access to user email'),
('api:read', 'Read-only API access'),
('api:write', 'Write API access');