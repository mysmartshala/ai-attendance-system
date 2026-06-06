<?php
require_once '../../config.php';
require_once '../../includes/api_client.php';

if (!isset($_SESSION['token'])) {
    header('Location: /auth/login.php');
    exit;
}

$api = new APIClient();
$dashboard = $api->getDashboard();

$page_title = 'Teacher Dashboard';
?>

<?php include '../../includes/header.php'; ?>

<div class="row mt-4">
    <div class="col-md-3">
        <div class="card text-white bg-primary">
            <div class="card-body">
                <h6 class="card-title">Total Students</h6>
                <h2><?php echo $dashboard['total_students'] ?? 0; ?></h2>
            </div>
        </div>
    </div>

    <div class="col-md-3">
        <div class="card text-white bg-success">
            <div class="card-body">
                <h6 class="card-title">Today's Present</h6>
                <h2><?php echo $dashboard['todays_attendance'] ?? 0; ?></h2>
            </div>
        </div>
    </div>

    <div class="col-md-3">
        <div class="card text-white bg-danger">
            <div class="card-body">
                <h6 class="card-title">Today's Absent</h6>
                <h2><?php echo $dashboard['todays_absence'] ?? 0; ?></h2>
            </div>
        </div>
    </div>

    <div class="col-md-3">
        <div class="card text-white bg-info">
            <div class="card-body">
                <h6 class="card-title">Attendance %</h6>
                <h2><?php echo round($dashboard['attendance_percentage'] ?? 0, 1); ?>%</h2>
            </div>
        </div>
    </div>
</div>

<div class="row mt-4">
    <div class="col-md-12">
        <div class="card">
            <div class="card-header">
                <h5 class="mb-0">Quick Actions</h5>
            </div>
            <div class="card-body">
                <a href="/teacher/attendance.php" class="btn btn-primary btn-lg me-2">
                    <i class="fas fa-camera"></i> Mark Attendance
                </a>
                <a href="/teacher/analytics.php" class="btn btn-info btn-lg me-2">
                    <i class="fas fa-chart-bar"></i> View Analytics
                </a>
            </div>
        </div>
    </div>
</div>

<?php include '../../includes/footer.php'; ?>