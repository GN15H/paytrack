CREATE TABLE categories (
  id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
  user_id UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE,
  name VARCHAR(255) NOT NULL,
  type VARCHAR(50) NOT NULL CHECK (type IN ('income', 'expense')),
  created_at TIMESTAMP NOT NULL DEFAULT NOW()
);

-- INSERT INTO categories (id, user_id, name, type) VALUES
--   (uuid_generate_v4(), '00000000-0000-0000-0000-000000000000', 'Salary', 'income'),
--   (uuid_generate_v4(), '00000000-0000-0000-0000-000000000000', 'Food', 'expense'),
--   (uuid_generate_v4(), '00000000-0000-0000-0000-000000000000', 'Transport', 'expense'),
--   (uuid_generate_v4(), '00000000-0000-0000-0000-000000000000', 'Entertainment', 'expense');
