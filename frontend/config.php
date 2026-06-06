<?php
define('API_BASE_URL', 'http://localhost:8080/api');
define('SESSION_TIMEOUT', 3600);
define('UPLOAD_MAX_SIZE', 50 * 1024 * 1024);
define('ALLOWED_EXTENSIONS', ['jpg', 'jpeg', 'png', 'gif']);

define('DB_HOST', 'localhost');
define('DB_USER', 'root');
define('DB_PASS', 'root');
define('DB_NAME', 'attendance');

session_start();

if (isset($_SESSION['last_activity']) && time() - $_SESSION['last_activity'] > SESSION_TIMEOUT) {
    session_destroy();
    header('Location: /auth/login.php');
    exit;
}

$_SESSION['last_activity'] = time();
?>