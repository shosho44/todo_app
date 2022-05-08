-- DBサーバーで初期のセットアップ用sql

DROP DATABASE IF EXISTS todo_app;

CREATE DATABASE todo_app;
USE todo_app;

DROP TABLE IF EXISTS tasks, task_categories, task_repeat, task_status;

CREATE TABLE task_categories (
    id INT AUTO_INCREMENT NOT NULL PRIMARY KEY,
    category_name VARCHAR(255) NOT NULL,
    created_at DATETIME NOT NULL DEFAULT NOW(),
    updated_at DATETIME NOT NULL DEFAULT NOW(),
    INDEX (id)
) ENGINE=INNODB COMMENT='アプリケーション側で、クライアントからid=1の削除はできないようにしている。tasksからの参照をしているカラムのデフォルトが1のため。'
;

CREATE TABLE task_repeat (
    id INT AUTO_INCREMENT NOT NULL PRIMARY KEY,
    repeat_type VARCHAR(255) NOT NULL,
    created_at DATETIME NOT NULL DEFAULT NOW(),
    updated_at DATETIME NOT NULL DEFAULT NOW()
) ENGINE=INNODB COMMENT='アプリケーション側で、クライアントからid=1の削除はできないようにしている。tasksからの参照をしているカラムのデフォルトが1のため。'
;

CREATE TABLE task_status (
    id INT AUTO_INCREMENT NOT NULL PRIMARY KEY,
    status_type VARCHAR(255) NOT NULL,
    created_at DATETIME NOT NULL DEFAULT NOW(),
    updated_at DATETIME NOT NULL DEFAULT NOW()
) ENGINE=INNODB COMMENT='アプリケーション側で、クライアントからid=1の削除はできないようにしている。tasksからの参照をしているカラムのデフォルトが1のため。'
;

CREATE TABLE tasks (
    id INT AUTO_INCREMENT NOT NULL PRIMARY KEY,
    task_name VARCHAR(255) NOT NULL,
    deadline DATETIME NOT NULL DEFAULT '9999-12-31 00:00:00',
    category_id INT NOT NULL DEFAULT 1, -- foreign
    priority_percent_number INT NOT NULL DEFAULT 0,
    repeat_id INT NOT NULL DEFAULT 1, -- foreign
    status_id INT NOT NULL DEFAULT 1, -- foreign
    created_at DATETIME NOT NULL DEFAULT NOW(),
    updated_at DATETIME NOT NULL DEFAULT NOW(),
    INDEX (repeat_id),
    INDEX (status_id),
    INDEX (category_id),
    FOREIGN KEY (category_id)
        REFERENCES task_categories(id)
        ON UPDATE CASCADE
        ON DELETE RESTRICT,
    FOREIGN KEY (repeat_id)
        REFERENCES task_repeat(id)
        ON UPDATE CASCADE
        ON DELETE RESTRICT,
    FOREIGN KEY (status_id)
        REFERENCES task_status(id)
        ON UPDATE CASCADE
        ON DELETE RESTRICT
) ENGINE=INNODB
;