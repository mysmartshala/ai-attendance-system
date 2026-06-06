<?php
require_once '../../config.php';
require_once '../../includes/api_client.php';

$error = '';

if ($_SERVER['REQUEST_METHOD'] === 'POST') {
    $username = $_POST['username'] ?? '';
    $password = $_POST['password'] ?? '';

    $api = new APIClient();
    $response = $api->login($username, $password);

    if (isset($response['token'])) {
        $_SESSION['token'] = $response['token'];
        $_SESSION['user_id'] = $response['teacher']['id'];
        $_SESSION['username'] = $response['teacher']['username'];
        header('Location: /teacher/dashboard.php');
        exit;
    } else {
        $error = 'Invalid username or password';
    }
}

$page_title = 'Teacher Login';
?>

<?php include '../../includes/header.php'; ?>

<div class="row justify-content-center mt-5">
    <div class="col-md-5">
        <div class="card shadow">
            <div class="card-body p-5">
                <h2 class="text-center mb-4">Teacher Login</h2>

                <?php if ($error): ?>
                    <div class="alert alert-danger">
                        <?php echo htmlspecialchars($error); ?>
                    </div>
                <?php endif; ?>

                <form method="POST">
                    <div class="mb-3">
                        <label for="username" class="form-label">Username</label>
                        <input type="text" class="form-control" id="username" name="username" required>
                    </div>

                    <div class="mb-3">
                        <label for="password" class="form-label">Password</label>
                        <input type="password" class="form-control" id="password" name="password" required>
                    </div>

                    <button type="submit" class="btn btn-primary w-100">Login</button>
                </form>

                <hr>
                <p class="text-muted small mt-3">
                    <strong>Demo Credentials:</strong><br>
                    Username: <code>teacher</code><br>
                    Password: <code>teacher123</code>
                </p>
            </div>
        </div>
    </div>
</div>

<?php include '../../includes/footer.php'; ?>