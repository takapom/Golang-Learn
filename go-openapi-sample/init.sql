-- データベースを作成（docker-compose.ymlで作成されるが念のため）
CREATE DATABASE IF NOT EXISTS todoapi_db;
USE todoapi_db;

-- todosテーブルを作成
CREATE TABLE IF NOT EXISTS todos (
    id BIGINT AUTO_INCREMENT PRIMARY KEY,
    title VARCHAR(255) NOT NULL,
    done BOOLEAN DEFAULT FALSE,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);

-- サンプルデータを挿入
INSERT INTO todos (title, done) VALUES
    ('Dockerセットアップを完了する', true),
    ('APIドキュメントを作成する', false),
    ('テストを書く', false);