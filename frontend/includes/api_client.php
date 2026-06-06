<?php

class APIClient {
    private $base_url;
    private $token;

    public function __construct($base_url = API_BASE_URL) {
        $this->base_url = $base_url;
        $this->token = $_SESSION['token'] ?? null;
    }

    public function login($username, $password) {
        return $this->post('/auth/login', [
            'username' => $username,
            'password' => $password,
        ]);
    }

    public function createStudent($data, $photoFile) {
        return $this->postMultipart('/students', $data, $photoFile, 'photo');
    }

    public function getStudents($filters = []) {
        $query = '';
        if (!empty($filters)) {
            $query = '?' . http_build_query($filters);
        }
        return $this->get('/students' . $query);
    }

    public function getStudent($id) {
        return $this->get('/students/' . $id);
    }

    public function updateStudent($id, $data, $photoFile = null) {
        if ($photoFile) {
            return $this->postMultipart('/students/' . $id, $data, $photoFile, 'photo', 'PUT');
        }
        return $this->put('/students/' . $id, $data);
    }

    public function deleteStudent($id) {
        return $this->delete('/students/' . $id);
    }

    public function processAttendance($classroomPhoto, $course, $semester, $date = null) {
        $data = [
            'course' => $course,
            'semester' => $semester,
        ];
        if ($date) {
            $data['date'] = $date;
        }
        return $this->postMultipart('/attendance/process', $data, $classroomPhoto, 'classroom_photo');
    }

    public function getAttendanceReport($course, $semester, $startDate = null, $endDate = null) {
        $params = [
            'course' => $course,
            'semester' => $semester,
        ];
        if ($startDate) {
            $params['start_date'] = $startDate;
        }
        if ($endDate) {
            $params['end_date'] = $endDate;
        }
        return $this->get('/attendance/report?' . http_build_query($params));
    }

    public function getDashboard() {
        return $this->get('/analytics/dashboard');
    }

    public function getCourseWiseAttendance() {
        return $this->get('/analytics/course-wise');
    }

    public function getStudentAnalytics($studentId) {
        return $this->get('/analytics/student-wise/' . $studentId);
    }

    private function get($endpoint) {
        $ch = curl_init($this->base_url . $endpoint);
        curl_setopt($ch, CURLOPT_RETURNTRANSFER, true);
        curl_setopt($ch, CURLOPT_HTTPHEADER, [
            'Authorization: Bearer ' . $this->token,
            'Content-Type: application/json',
        ]);
        $response = curl_exec($ch);
        curl_close($ch);
        return json_decode($response, true);
    }

    private function post($endpoint, $data) {
        $ch = curl_init($this->base_url . $endpoint);
        curl_setopt($ch, CURLOPT_RETURNTRANSFER, true);
        curl_setopt($ch, CURLOPT_POST, true);
        curl_setopt($ch, CURLOPT_HTTPHEADER, [
            'Authorization: Bearer ' . $this->token,
            'Content-Type: application/json',
        ]);
        curl_setopt($ch, CURLOPT_POSTFIELDS, json_encode($data));
        $response = curl_exec($ch);
        curl_close($ch);
        return json_decode($response, true);
    }

    private function put($endpoint, $data) {
        $ch = curl_init($this->base_url . $endpoint);
        curl_setopt($ch, CURLOPT_RETURNTRANSFER, true);
        curl_setopt($ch, CURLOPT_CUSTOMREQUEST, 'PUT');
        curl_setopt($ch, CURLOPT_HTTPHEADER, [
            'Authorization: Bearer ' . $this->token,
            'Content-Type: application/json',
        ]);
        curl_setopt($ch, CURLOPT_POSTFIELDS, json_encode($data));
        $response = curl_exec($ch);
        curl_close($ch);
        return json_decode($response, true);
    }

    private function delete($endpoint) {
        $ch = curl_init($this->base_url . $endpoint);
        curl_setopt($ch, CURLOPT_RETURNTRANSFER, true);
        curl_setopt($ch, CURLOPT_CUSTOMREQUEST, 'DELETE');
        curl_setopt($ch, CURLOPT_HTTPHEADER, [
            'Authorization: Bearer ' . $this->token,
        ]);
        $response = curl_exec($ch);
        curl_close($ch);
        return json_decode($response, true);
    }

    private function postMultipart($endpoint, $data, $file, $fileKey, $method = 'POST') {
        $ch = curl_init($this->base_url . $endpoint);
        
        $postData = $data;
        if (is_array($file) && isset($file['tmp_name'])) {
            $postData[$fileKey] = new CURLFile($file['tmp_name'], $file['type'], $file['name']);
        } else {
            $postData[$fileKey] = new CURLFile($file);
        }

        curl_setopt($ch, CURLOPT_RETURNTRANSFER, true);
        if ($method === 'PUT') {
            curl_setopt($ch, CURLOPT_CUSTOMREQUEST, 'PUT');
        } else {
            curl_setopt($ch, CURLOPT_POST, true);
        }
        curl_setopt($ch, CURLOPT_HTTPHEADER, [
            'Authorization: Bearer ' . $this->token,
        ]);
        curl_setopt($ch, CURLOPT_POSTFIELDS, $postData);
        $response = curl_exec($ch);
        curl_close($ch);
        return json_decode($response, true);
    }
}

?>