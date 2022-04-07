package com.slack.interviews;

import org.json.JSONObject;

public class Entity {

    private String string;
    private int integer;

    public Entity( String string, int integer) {
        this.string = string;
        this.integer = integer;
    }

    public static Entity fromJSON(JSONObject jsonObject) {
        return new Entity(
                jsonObject.optString("string", "unknown"),
                jsonObject.optInt("integer", -1));
    }

    public String getString() {
        return this.string;
    }

    public int getInteger() {
        return this.integer;
    }
}
