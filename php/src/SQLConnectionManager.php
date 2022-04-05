<?php

namespace BackendInterviewPrerequisites;

use SQLite3; 

final class SQLConnectionManager {
    const FILE_PATH = __DIR__.'/../datastore.db';
    const SCHEMA_PATH = __DIR__.'/../schema.sql';

    private static $sqlite;

    public static function getSQLConnection(): SQLite3 {
        if (static::$sqlite === null) {
            static::init();
        }
        return static::$sqlite;
    }

    public static function init() {
        static::$sqlite = new SQLite3(self::FILE_PATH);
        $schema_create = file_get_contents(static::SCHEMA_PATH);
        $ret = static::$sqlite->exec($schema_create);
        if (!$ret) {
            throw new Exception("SQL init failed: ".static::$sqlite->lastErrorMsg());
        }
    }

    public static function cleanup() {
        static::$sqlite->close();
        unlink(static::FILE_PATH);
    }
}