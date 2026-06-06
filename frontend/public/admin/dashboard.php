<?php
require_once '../../config.php';
require_once '../../includes/api_client.php';

if (!isset($_SESSION['token'])) {
    header('Location: /auth/login.php');
    exit;
}

$api = new APIClient();
$dashboard = $api->getDashboard();

$page_title = 'Admin Dashboard';
?>

<?php include '../../includes/header.php'; ?>

<div class="row mt-4">
    <div class="col-md-12">
        <h2>Admin Dashboard</h2>
    </div>
</div>

<div class="row mt-4">
    <div class="col-md-4">
        <div class="card">
            <div class="card-header">Quick Links</div>
            <div class="card-body">
                <ul class="list-group list-group-flush">
                    <li class="list-group-item">
                        <a href="/admin/students.php">Manage Students</a>
                    </li>
                    <li class="list-group-item">
                        <a href="/admin/add_student.php">Add New Student</a>
                    </li>
                </ul>
            </div>
        </div>
    </div>

    <div class="col-md-8">
        <div class="card">
            <div class="card-header">System Statistics</div>
            <div class="card-body">
                <table class="table">
                    <tr>
                        <td>Total Students</td>
                        <td><strong><?php echo $dashboard['total_students'] ?? 0; ?></strong></td>
                    </tr>
                    <tr>
                        <td>Today's Attendance</td>
                        <td><strong><?php echo $dashboard['todays_attendance'] ?? 0; ?></strong></td>
                    </tr>
                </table>
            </div>
        </div>
    </div>
</div>

<?php include '../../includes/footer.php'; ?>