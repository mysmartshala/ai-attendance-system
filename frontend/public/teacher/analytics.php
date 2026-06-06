<?php
require_once '../../config.php';
require_once '../../includes/api_client.php';

if (!isset($_SESSION['token'])) {
    header('Location: /auth/login.php');
    exit;
}

$api = new APIClient();
$courseWise = $api->getCourseWiseAttendance();

$page_title = 'Analytics Dashboard';
?>

<?php include '../../includes/header.php'; ?>

<div class="row mt-4">
    <div class="col-md-12">
        <div class="card">
            <div class="card-header">
                <h5 class="mb-0">Course-Wise Attendance</h5>
            </div>
            <div class="card-body">
                <canvas id="courseChart" height="80"></canvas>
            </div>
        </div>
    </div>
</div>

<div class="row mt-4">
    <div class="col-md-12">
        <div class="card">
            <div class="card-header">
                <h5 class="mb-0">Attendance Report</h5>
            </div>
            <div class="card-body">
                <div class="row mb-3">
                    <div class="col-md-3">
                        <select class="form-select" id="filter_course">
                            <option value="">Select Course</option>
                            <option value="BCA">BCA</option>
                            <option value="BCom">BCom</option>
                            <option value="BBA">BBA</option>
                        </select>
                    </div>
                    <div class="col-md-3">
                        <select class="form-select" id="filter_semester">
                            <option value="">Select Semester</option>
                            <option value="1">1</option>
                            <option value="2">2</option>
                            <option value="3">3</option>
                            <option value="4">4</option>
                            <option value="5">5</option>
                            <option value="6">6</option>
                        </select>
                    </div>
                    <div class="col-md-3">
                        <input type="date" id="filter_start_date" class="form-control">
                    </div>
                    <div class="col-md-3">
                        <button class="btn btn-primary w-100" onclick="loadReport()">Filter</button>
                    </div>
                </div>
                <div id="reportTable"></div>
            </div>
        </div>
    </div>
</div>

<script>
    const courseData = <?php echo json_encode($courseWise ?? []); ?>;
    const courses = courseData.map(d => d.Course || 'Unknown');
    const percentages = courseData.map(d => d.Percentage || 0);

    const ctx = document.getElementById('courseChart').getContext('2d');
    new Chart(ctx, {
        type: 'bar',
        data: {
            labels: courses,
            datasets: [{
                label: 'Attendance %',
                data: percentages,
                backgroundColor: 'rgba(75, 192, 192, 0.5)',
                borderColor: 'rgba(75, 192, 192, 1)',
                borderWidth: 1
            }]
        },
        options: {
            responsive: true,
            scales: {
                y: {
                    beginAtZero: true,
                    max: 100
                }
            }
        }
    });

    function loadReport() {
        const course = document.getElementById('filter_course').value;
        const semester = document.getElementById('filter_semester').value;

        if (!course || !semester) {
            alert('Please select course and semester');
            return;
        }

        fetch('http://localhost:8080/api/attendance/report?' + new URLSearchParams({
            course: course,
            semester: semester
        }), {
            headers: {
                'Authorization': 'Bearer ' + '<?php echo $_SESSION['token']; ?>'
            }
        })
        .then(r => r.json())
        .then(data => {
            let html = '<table class="table table-striped"><thead><tr>';
            html += '<th>Roll No</th><th>Name</th><th>Present</th><th>Absent</th><th>%</th></tr></thead><tbody>';
            data.forEach(row => {
                html += `<tr><td>${row.roll_no}</td><td>${row.name}</td>`;
                html += `<td>${row.present}</td><td>${row.absent}</td><td>${row.percentage}%</td></tr>`;
            });
            html += '</tbody></table>';
            document.getElementById('reportTable').innerHTML = html;
        });
    }
</script>

<?php include '../../includes/footer.php'; ?>