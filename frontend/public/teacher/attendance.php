<?php
require_once '../../config.php';
require_once '../../includes/api_client.php';

if (!isset($_SESSION['token'])) {
    header('Location: /auth/login.php');
    exit;
}

$api = new APIClient();
$result = null;
$error = '';

if ($_SERVER['REQUEST_METHOD'] === 'POST') {
    $course = $_POST['course'] ?? '';
    $semester = $_POST['semester'] ?? '';
    $date = $_POST['date'] ?? date('Y-m-d');

    if ($_FILES['classroom_photo']['error'] === UPLOAD_ERR_OK) {
        $result = $api->processAttendance($_FILES['classroom_photo'], $course, $semester, $date);
        if (!$result || isset($result['error'])) {
            $error = $result['error'] ?? 'Failed to process attendance';
        }
    } else {
        $error = 'Please upload a classroom photo';
    }
}

$page_title = 'Mark Attendance';
?>

<?php include '../../includes/header.php'; ?>

<div class="row mt-4">
    <div class="col-md-6">
        <div class="card">
            <div class="card-header">
                <h5 class="mb-0">Upload Classroom Photo</h5>
            </div>
            <div class="card-body">
                <?php if ($error): ?>
                    <div class="alert alert-danger"><?php echo htmlspecialchars($error); ?></div>
                <?php endif; ?>

                <form method="POST" enctype="multipart/form-data">
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
                        <label for="date" class="form-label">Date</label>
                        <input type="date" class="form-control" id="date" name="date" value="<?php echo date('Y-m-d'); ?>">
                    </div>

                    <div class="mb-3">
                        <label for="classroom_photo" class="form-label">Classroom Photo</label>
                        <input type="file" class="form-control" id="classroom_photo" name="classroom_photo" accept="image/*" required>
                        <small class="text-muted">You can also use camera on mobile</small>
                    </div>

                    <button type="submit" class="btn btn-primary w-100">Process Attendance</button>
                </form>
            </div>
        </div>
    </div>

    <div class="col-md-6">
        <?php if ($result): ?>
            <div class="card">
                <div class="card-header bg-primary text-white">
                    <h5 class="mb-0">Attendance Results</h5>
                </div>
                <div class="card-body">
                    <table class="table table-sm">
                        <tr>
                            <td><strong>Total Students:</strong></td>
                            <td><?php echo $result['total_students']; ?></td>
                        </tr>
                        <tr>
                            <td><strong>Faces Detected:</strong></td>
                            <td><?php echo $result['detected']; ?></td>
                        </tr>
                        <tr>
                            <td><strong>Present:</strong></td>
                            <td><span class="badge bg-success"><?php echo $result['present']; ?></span></td>
                        </tr>
                        <tr>
                            <td><strong>Unknown:</strong></td>
                            <td><span class="badge bg-warning"><?php echo $result['unknown']; ?></span></td>
                        </tr>
                        <tr>
                            <td><strong>Absent:</strong></td>
                            <td><span class="badge bg-danger"><?php echo $result['absent']; ?></span></td>
                        </tr>
                    </table>

                    <h6 class="mt-4">Matched Students</h6>
                    <div class="table-responsive">
                        <table class="table table-sm table-striped">
                            <thead>
                                <tr>
                                    <th>Name</th>
                                    <th>Roll No</th>
                                    <th>Confidence</th>
                                    <th>Status</th>
                                </tr>
                            </thead>
                            <tbody>
                                <?php foreach ($result['match_results'] as $match): ?>
                                    <?php if ($match['is_matched']): ?>
                                        <tr>
                                            <td><?php echo htmlspecialchars($match['student_name']); ?></td>
                                            <td><?php echo htmlspecialchars($match['roll_no']); ?></td>
                                            <td><?php echo round($match['confidence'] * 100, 1); ?>%</td>
                                            <td><span class="badge bg-success">✓</span></td>
                                        </tr>
                                    <?php endif; ?>
                                <?php endforeach; ?>
                            </tbody>
                        </table>
                    </div>

                    <h6 class="mt-4">Unknown Faces</h6>
                    <div class="table-responsive">
                        <table class="table table-sm table-striped">
                            <thead>
                                <tr>
                                    <th>Face Index</th>
                                    <th>Status</th>
                                </tr>
                            </thead>
                            <tbody>
                                <?php foreach ($result['match_results'] as $match): ?>
                                    <?php if (!$match['is_matched']): ?>
                                        <tr>
                                            <td><?php echo $match['face_index']; ?></td>
                                            <td><span class="badge bg-warning">?</span></td>
                                        </tr>
                                    <?php endif; ?>
                                <?php endforeach; ?>
                            </tbody>
                        </table>
                    </div>
                </div>
            </div>
        <?php endif; ?>
    </div>
</div>

<?php include '../../includes/footer.php'; ?>