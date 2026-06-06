<?php
require_once '../../config.php';
require_once '../../includes/api_client.php';

if (!isset($_SESSION['token'])) {
    header('Location: /auth/login.php');
    exit;
}

$error = '';
$success = '';

if ($_SERVER['REQUEST_METHOD'] === 'POST') {
    $api = new APIClient();
    $result = $api->createStudent($_POST, $_FILES['photo']);

    if (isset($result['id'])) {
        $success = 'Student added successfully!';
        header('Location: /admin/students.php');
        exit;
    } else {
        $error = $result['error'] ?? 'Failed to add student';
    }
}

$page_title = 'Add Student';
?>

<?php include '../../includes/header.php'; ?>

<div class="row mt-4 justify-content-center">
    <div class="col-md-6">
        <div class="card">
            <div class="card-header">
                <h5 class="mb-0">Add New Student</h5>
            </div>
            <div class="card-body">
                <?php if ($error): ?>
                    <div class="alert alert-danger"><?php echo htmlspecialchars($error); ?></div>
                <?php endif; ?>

                <form method="POST" enctype="multipart/form-data">
                    <div class="mb-3">
                        <label for="roll_no" class="form-label">Roll Number</label>
                        <input type="text" class="form-control" id="roll_no" name="roll_no" required>
                    </div>

                    <div class="mb-3">
                        <label for="name" class="form-label">Name</label>
                        <input type="text" class="form-control" id="name" name="name" required>
                    </div>

                    <div class="mb-3">
                        <label for="course" class="form-label">Course</label>
                        <select class="form-select" id="course" name="course" required>
                            <option value="">Select Course</option>
                            <option value="BCA">BCA</option>
                            <option value="BCom">BCom</option>
                            <option value="BBA">BBA</option>
                        </select>
                    </div>

                    <div class="mb-3">
                        <label for="semester" class="form-label">Semester</label>
                        <select class="form-select" id="semester" name="semester" required>
                            <option value="">Select Semester</option>
                            <option value="1">1</option>
                            <option value="2">2</option>
                            <option value="3">3</option>
                            <option value="4">4</option>
                            <option value="5">5</option>
                            <option value="6">6</option>
                        </select>
                    </div>

                    <div class="mb-3">
                        <label for="photo" class="form-label">Photo</label>
                        <input type="file" class="form-control" id="photo" name="photo" accept="image/*" required>
                    </div>

                    <button type="submit" class="btn btn-primary w-100">Add Student</button>
                    <a href="/admin/students.php" class="btn btn-secondary w-100 mt-2">Cancel</a>
                </form>
            </div>
        </div>
    </div>
</div>

<?php include '../../includes/footer.php'; ?>