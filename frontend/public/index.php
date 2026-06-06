<?php
require_once '../config.php';

if (isset($_SESSION['token'])) {
    header('Location: /teacher/dashboard.php');
    exit;
} else {
    header('Location: /auth/login.php');
    exit;
}
?>