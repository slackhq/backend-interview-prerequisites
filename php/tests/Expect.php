<?php

function expect($actual): ExpectObj {
    return new ExpectObj($actual);
}

final class ExpectObj {
    private $actual;
    public static $failures = 0;

    public function __construct($actual) {
        $this->actual = $actual;
    }

    public function toBeSame($expected, string $description) {
        if ($this->actual === $expected) {
            $this->pass($description);
        } else {
            $description = "$description".PHP_EOL."Expected: ".$expected.PHP_EOL."Got: ".$this->actual.PHP_EOL;
            $this->fail($description);
        }
    }

    public function toBeTruthy(string $description) {
        if ((bool)$this->actual) {
            $this->pass($description);
        } else {
            $this->fail($description);
        }
    }

    private function fail(string $description) {
        $line = debug_backtrace(DEBUG_BACKTRACE_IGNORE_ARGS)[1]['line'];
        echo "❌ Failure: $description (line $line)".PHP_EOL;
        static::$failures++;
    }

    private function pass(string $description) {
        $line = debug_backtrace(DEBUG_BACKTRACE_IGNORE_ARGS)[1]['line'];
        echo "✅ OK: $description (line $line)".PHP_EOL;
    }
}