document.addEventListener('DOMContentLoaded', function() {
    console.log('AI Attendance System loaded');
});

function deleteStudent(id) {
    if (confirm('Are you sure you want to delete this student?')) {
        fetch('http://localhost:8080/api/students/' + id, {
            method: 'DELETE',
            headers: {
                'Authorization': 'Bearer ' + localStorage.getItem('token')
            }
        })
        .then(response => {
            if (response.ok) {
                alert('Student deleted successfully');
                location.reload();
            } else {
                alert('Failed to delete student');
            }
        });
    }
}

function storeToken(token) {
    localStorage.setItem('token', token);
}

function getToken() {
    return localStorage.getItem('token');
}

function clearToken() {
    localStorage.removeItem('token');
}
