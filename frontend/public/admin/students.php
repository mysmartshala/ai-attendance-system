<?php
require_once '../../config.php';
require_once '../../includes/api_client.php';

if (!isset($_SESSION['token'])) {
    header('Location: /auth/login.php');
    exit;
}

$api = new APIClient();
$students = $api->getStudents();

$page_title = 'Student List';
?>

<?php include '../../includes/header.php'; ?>

<div class="row mt-4">
    <div class="col-md-12">
        <div class="card">
            <div class="card-header">
                <h5 class="mb-0 d-flex justify-content-between align-items-center">
                    <span>Students</span>
                    <a href="/admin/add_student.php" class="btn btn-sm btn-primary">Add Student</a>
                </h5>
            </div>
            <div class="card-body">
                <div class="table-responsive">
                    <table class="table table-striped">
                        <thead>
                            <tr>
                                <th>Roll No</th>
                                <th>Name</th>
                                <th>Course</th>
                                <th>Semester</th>
                                <th>Actions</th>
                            </tr>
                        </thead>
                        <tbody>
                            <?php if (is_array($students)): ?>
                                <?php foreach ($students as $student): ?>
                                    <tr>
                                        <td><?php echo htmlspecialchars($student['roll_no'] ?? ''); ?></td>
                                        <td><?php echo htmlspecialchars($student['name'] ?? ''); ?></td>
                                        <td><?php echo htmlspecialchars($student['course'] ?? ''); ?></td>
                                        <td><?php echo $student['semester'] ?? ''; ?></td>
                                        <td>
                                            <a href="/admin/edit_student.php?id=<?php echo $student['id']; ?>" class="btn btn-sm btn-warning">Edit</a>
                                            <button class="btn btn-sm btn-danger" onclick="deleteStudent(<?php echo $student['id']; ?>)">Delete</button>
                                        </td>
                                    </tr>
                                <?php endforeach; ?>
                            <?php endif; ?>
                        </tbody>
                    </table>
                </div>
            </div>
        </div>
    </div>
</div>

<?php include '../../includes/footer.php'; ?>