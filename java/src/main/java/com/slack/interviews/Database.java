
package com.slack.interviews;

import java.io.IOException;
import java.sql.Connection;
import java.sql.DriverManager;
import java.sql.PreparedStatement;
import java.sql.ResultSet;
import java.sql.Statement;
import java.sql.SQLException;

import java.util.Optional;

public class Database {

    private Connection connection;

    public Database(String databaseUrl) throws ClassNotFoundException, SQLException, IOException {
        Class.forName("org.sqlite.JDBC");
        connection = DriverManager.getConnection(databaseUrl);
        load(connection);
    }

    private void load(Connection connection) throws SQLException, IOException {
        connection.createStatement().executeUpdate(
                "CREATE TABLE IF NOT EXISTS test_table ( \n" +
                        "   id INTEGER NOT NULL PRIMARY KEY, \n" +
                        "   txt TEXT NOT NULL \n" +
                        "); \n");
    }

    public Optional<String> getById(long id) throws SQLException {
        PreparedStatement statement = this.connection.prepareStatement(
                "SELECT * \n"
                        + "FROM \n"
                        + "   test_table \n"
                        + "WHERE \n"
                        + "   id = ?\n"
                        + ";");
        statement.setLong(1, id);
        ResultSet resultSet = statement.executeQuery();
        if (!resultSet.next()) {
            return Optional.empty();
        }
        return Optional.of(resultSet.getString("txt"));
    }

    public Optional<Long> addRow(String txt) throws SQLException {
        PreparedStatement statement = connection.prepareStatement(
                "INSERT INTO test_table \n"
                        + " (id, txt) \n"
                        + " VALUES \n"
                        + " (NULL, ?);",
                Statement.RETURN_GENERATED_KEYS);

        statement.setString(1, txt);

        if (statement.executeUpdate() == 0) {
            return Optional.empty();
        }

        return Optional.of(statement.getGeneratedKeys().getLong(1));
    }
}
