//Функция который проверяет почту на правописание или на формат
function validateEmail(email) {
    const emailPattern = /^[a-zA-Z0-9._-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,6}$/;
    let errors = [];

    if (!emailPattern.test(email)) {
        errors.push('Please enter a valid email address.');
    }

    return errors;
}
//Функция который проверяет Пароль
function validatePassword(password) {
    // Минимальная длина 8 символов.
    // Наличие хотя бы одной заглавной буквы.
    // Наличие хотя бы одной строчной буквы.
    // Наличие хотя бы одной цифры.
        const minLength = 8;
        const hasUpperCase = /[A-Z]/.test(password);
        const hasLowerCase = /[a-z]/.test(password);
        const hasDigit = /\d/.test(password);
    
        let errors = [];
    
        if (password.length < minLength) {
            errors.push(`Password must be at least ${minLength} characters long.`);
            return errors
        }
        if (!hasUpperCase) {
            errors.push('Password must contain at least one uppercase letter.');
            return errors
        }
        if (!hasLowerCase) {
            errors.push('Password must contain at least one lowercase letter.');
            return errors
        }
        if (!hasDigit) {
            errors.push('Password must contain at least one digit.');
            return errors
        }
    
        return errors;
}


//Реагирует на нажатие сабмит
document.getElementById('signInForm').addEventListener('submit', function(event) {
    event.preventDefault();
    const email = document.getElementById('email').value;
    const password = document.getElementById('password').value;

    const emailErrors = validateEmail(email);
    const passwordErrors = validatePassword(password);
    const errorMessage = document.getElementById('errorMessage');

    if (emailErrors.length > 0 || passwordErrors.length > 0) {
        if (emailErrors.length > 0) {
            errorMessage.innerHTML = emailErrors.join('<br>');
            return
        }
        if (passwordErrors.length > 0) {
            errorMessage.innerHTML = passwordErrors.join('<br>');
            return
        }
    }
    const data = {Email : email, Password : password};  // Создаем объект с данными
    const url = 'http://localhost:8080/api/signIn';  // URL-адрес сервера
        
        // Отправляем запрос на сервер
        fetch(url, {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json'
            },
            body: JSON.stringify(data)
        })
        .then(response => {
            if (!response.ok){
                errorMessage.innerHTML = "Login or password is incorrect"
                return
            } else{
                console.log("okk")
                
                // Перенаправление на другую страницу
                window.location.href = "http://localhost:8080/";
            }
        })  // Обрабатываем ответ как JSON

        .catch(error => console.error('Error:', error));  // Обрабатываем возможные ошибки
    
});