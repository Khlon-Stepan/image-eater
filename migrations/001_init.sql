CREATE TABLE image_tasks (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    original_filename VARCHAR(255) NOT NULL,
    -- Возможные статусы: pending (в очереди), processing (в работе), done (готово), failed (ошибка)
    status VARCHAR(50) NOT NULL DEFAULT 'pending',
    
    -- Ссылки на файлы в MinIO (появятся после загрузки/обработки)
    original_url TEXT,
    processed_url TEXT,
    
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

-- Индекс для быстрого поиска задач по статусу (например, чтобы найти зависшие)
CREATE INDEX idx_image_tasks_status ON image_tasks(status);