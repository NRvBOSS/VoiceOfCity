-- +goose Up
CREATE TABLE issues (
    id BIGINT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    title VARCHAR(150) NOT NULL,
    description TEXT NOT NULL DEFAULT '',
    category VARCHAR(50) NOT NULL,
    latitude DOUBLE PRECISION NOT NULL,
    longitude DOUBLE PRECISION NOT NULL,
    status VARCHAR(20) NOT NULL DEFAULT 'open',
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),

    CONSTRAINT issues_category_check
        CHECK (category IN ('road', 'waste', 'lighting', 'sidewalk', 'other')),

    CONSTRAINT issues_status_check
        CHECK (status IN ('open', 'in_progress', 'resolved')),

    CONSTRAINT issues_latitude_check
        CHECK (latitude BETWEEN -90 AND 90),

    CONSTRAINT issues_longitude_check
        CHECK (longitude BETWEEN -180 AND 180)
);

CREATE INDEX idx_issues_created_at ON issues (created_at DESC);
CREATE INDEX idx_issues_status ON issues (status);

-- +goose Down
DROP TABLE issues;