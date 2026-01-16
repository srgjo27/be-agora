-- ============================================
-- Agora Forum Database Schema
-- ============================================

-- Enable UUID extension
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

-- ============================================
-- Table: users
-- ============================================
CREATE TABLE IF NOT EXISTS users (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    username VARCHAR(50) NOT NULL UNIQUE,
    email VARCHAR(255) NOT NULL UNIQUE,
    password_hash VARCHAR(255) NOT NULL,
    avatar_url TEXT,
    role VARCHAR(20) NOT NULL DEFAULT 'member' CHECK (role IN ('admin', 'moderator', 'member')),
    created_at TIMESTAMP NOT NULL DEFAULT NOW()
);

-- Indexes for users
CREATE INDEX idx_users_email ON users(email);
CREATE INDEX idx_users_username ON users(username);
CREATE INDEX idx_users_role ON users(role);

-- ============================================
-- Table: categories
-- ============================================
CREATE TABLE IF NOT EXISTS categories (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    name VARCHAR(100) NOT NULL UNIQUE,
    slug VARCHAR(100) NOT NULL UNIQUE,
    description TEXT,
    created_at TIMESTAMP NOT NULL DEFAULT NOW()
);

-- Indexes for categories
CREATE INDEX idx_categories_slug ON categories(slug);

-- ============================================
-- Table: threads
-- ============================================
CREATE TABLE IF NOT EXISTS threads (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    title VARCHAR(255) NOT NULL,
    slug VARCHAR(255) NOT NULL UNIQUE,
    content TEXT NOT NULL,
    user_id UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    category_id UUID NOT NULL REFERENCES categories(id) ON DELETE CASCADE,
    is_pinned BOOLEAN NOT NULL DEFAULT FALSE,
    is_locked BOOLEAN NOT NULL DEFAULT FALSE,
    vote_count INTEGER NOT NULL DEFAULT 0,
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP
);

-- Indexes for threads
CREATE INDEX idx_threads_user_id ON threads(user_id);
CREATE INDEX idx_threads_category_id ON threads(category_id);
CREATE INDEX idx_threads_slug ON threads(slug);
CREATE INDEX idx_threads_created_at ON threads(created_at DESC);
CREATE INDEX idx_threads_vote_count ON threads(vote_count DESC);
CREATE INDEX idx_threads_is_pinned ON threads(is_pinned);

-- ============================================
-- Table: posts
-- ============================================
CREATE TABLE IF NOT EXISTS posts (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    content TEXT NOT NULL,
    user_id UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    thread_id UUID NOT NULL REFERENCES threads(id) ON DELETE CASCADE,
    parent_post_id UUID REFERENCES posts(id) ON DELETE CASCADE,
    vote_count INTEGER NOT NULL DEFAULT 0,
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP
);

-- Indexes for posts
CREATE INDEX idx_posts_user_id ON posts(user_id);
CREATE INDEX idx_posts_thread_id ON posts(thread_id);
CREATE INDEX idx_posts_parent_post_id ON posts(parent_post_id);
CREATE INDEX idx_posts_created_at ON posts(created_at ASC);

-- ============================================
-- Table: thread_votes
-- ============================================
CREATE TABLE IF NOT EXISTS thread_votes (
    user_id UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    thread_id UUID NOT NULL REFERENCES threads(id) ON DELETE CASCADE,
    vote_type INTEGER NOT NULL CHECK (vote_type IN (-1, 0, 1)),
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    PRIMARY KEY (user_id, thread_id)
);

-- Indexes for thread_votes
CREATE INDEX idx_thread_votes_thread_id ON thread_votes(thread_id);

-- ============================================
-- Table: post_votes
-- ============================================
CREATE TABLE IF NOT EXISTS post_votes (
    user_id UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    post_id UUID NOT NULL REFERENCES posts(id) ON DELETE CASCADE,
    vote_type INTEGER NOT NULL CHECK (vote_type IN (-1, 0, 1)),
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    PRIMARY KEY (user_id, post_id)
);

-- Indexes for post_votes
CREATE INDEX idx_post_votes_post_id ON post_votes(post_id);

-- ============================================
-- Seed Data (Optional)
-- ============================================

-- Insert default admin user (password: admin123)
-- Password hash generated with bcrypt cost 10
INSERT INTO users (id, username, email, password_hash, role, created_at)
VALUES (
    uuid_generate_v4(),
    'admin',
    'admin@agora.com',
    '$2a$10$YourHashedPasswordHere', -- Replace with actual bcrypt hash
    'admin',
    NOW()
) ON CONFLICT (email) DO NOTHING;

-- Insert default categories
INSERT INTO categories (id, name, slug, description, created_at)
VALUES
    (uuid_generate_v4(), 'General Discussion', 'general-discussion', 'General topics and discussions', NOW()),
    (uuid_generate_v4(), 'Technology', 'technology', 'Tech news, programming, and development', NOW()),
    (uuid_generate_v4(), 'Off Topic', 'off-topic', 'Casual conversations and random topics', NOW())
ON CONFLICT (slug) DO NOTHING;

-- ============================================
-- Functions & Triggers (Optional)
-- ============================================

-- Function to update updated_at timestamp
CREATE OR REPLACE FUNCTION update_updated_at_column()
RETURNS TRIGGER AS $$
BEGIN
    NEW.updated_at = NOW();
    RETURN NEW;
END;
$$ language 'plpgsql';

-- Trigger for threads
CREATE TRIGGER update_threads_updated_at
    BEFORE UPDATE ON threads
    FOR EACH ROW
    EXECUTE FUNCTION update_updated_at_column();

-- Trigger for posts
CREATE TRIGGER update_posts_updated_at
    BEFORE UPDATE ON posts
    FOR EACH ROW
    EXECUTE FUNCTION update_updated_at_column();
