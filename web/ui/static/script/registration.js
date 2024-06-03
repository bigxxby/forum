
//Функция который проверяет логин
function validateUsername(username) {
    const minLength = 5;
    const isValid = /^[a-zA-Z0-9_]+$/.test(username); // Логин может содержать только буквы, цифры и подчеркивания
    let errors = [];

    if (username.length < minLength) {
        errors.push(`Username must be at least ${minLength} characters long.`);
        return errors
    }
    if (!isValid) {
        errors.push('Username can only contain letters, numbers, and underscores.');
        return errors
    }
    return errors;
}
//Функция который проверяет логин на совпадение
async function checkUserExists(username) {
    const url = `http://localhost:8080/api/users/taken?login=${encodeURIComponent(username)}`;

    try {
        const response = await fetch(url);
        if (!response.ok) {
            return false;
        }
    } catch (error) {
        console.error('Error:', error);
        return false; // Возвращаем false в случае ошибки
    }

    return true;
}

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
//Функция который проверяет два пароля
function validateConfirmPassword(password, confirmPassword) {
    let errors = [];

    if (password !== confirmPassword) {
        errors.push('Passwords do not match.');
    }

    return errors;
}
//Реагирует на заполнение логина
document.getElementById('username').addEventListener('input', function() {
    const username = this.value;
    const usernameError = document.getElementById('errorMessage');
    const errors = validateUsername(username);

    if (errors.length > 0) {
        usernameError.innerHTML = errors.join('<br>');
        this.setCustomValidity('Invalid field.');
    } else {
        usernameError.innerHTML = '';
        this.setCustomValidity('');
    }
});

document.getElementById('email').addEventListener('input', function() {
    const email = this.value;
    const emailError = document.getElementById('errorMessage');
    const errors = validateEmail(email);

    if (errors.length > 0) {
        emailError.innerHTML = errors.join('<br>');
        this.setCustomValidity('Invalid field.');
    } else {
        emailError.innerHTML = '';
        this.setCustomValidity('');
    }
});

document.getElementById('password').addEventListener('input', function() {
    const password = this.value;
    const passwordError = document.getElementById('errorMessage');
    const errors = validatePassword(password);

    if (errors.length > 0) {
        passwordError.innerHTML = errors.join('<br>');
        this.setCustomValidity('Invalid field.');
    } else {
        passwordError.innerHTML = '';
        this.setCustomValidity('');
    }
});

document.getElementById('confirmPassword').addEventListener('input', function() {
    const confirmPassword = this.value;
    const password = document.getElementById('password').value;
    const confirmPasswordError = document.getElementById('errorMessage');
    const errors = validateConfirmPassword(password, confirmPassword);

    if (errors.length > 0) {
        confirmPasswordError.innerHTML = errors.join('<br>');
        this.setCustomValidity('Invalid field.');
    } else {
        confirmPasswordError.innerHTML = '';
        this.setCustomValidity('');
    }
});
//Реагирует на нажатие сабмит
document.getElementById('authorizationForm').addEventListener('submit', function(event) {
    event.preventDefault();
    const username = document.getElementById('username').value;
    const email = document.getElementById('email').value;
    const password = document.getElementById('password').value;
    const confirmPassword = document.getElementById('confirmPassword').value;
    const errorMessage = document.getElementById('errorMessage');
    
    const usernameErrors = validateUsername(username);
    checkUserExists(username).then(repeatUsername => {
        errorMessage.innerHTML = confirmPasswordErrors.join('<br>');
        return
    });

    const emailErrors = validateEmail(email);
    const passwordErrors = validatePassword(password);
    const confirmPasswordErrors = validateConfirmPassword(password, confirmPassword);

    if (usernameErrors.length > 0 || emailErrors.length > 0 || passwordErrors.length > 0 || confirmPasswordErrors.length > 0) {
        if (usernameErrors.length > 0) {
            errorMessage.innerHTML = usernameErrors.join('<br>');
            return
        }
        if (emailErrors.length > 0) {
            errorMessage.innerHTML = emailErrors.join('<br>');
            return
        }
        if (passwordErrors.length > 0) {
            errorMessage.innerHTML = passwordErrors.join('<br>');
            return
        }
        if (confirmPasswordErrors.length > 0) {
            errorMessage.innerHTML = confirmPasswordErrors.join('<br>');
            return
        }
        
        if (repeatUsername.length > 0){
            errorMessage.innerHTML = confirmPasswordErrors.join('<br>');
            return
        }

    }
    const data = { Login : username, Email : email, Password : password, ConfirmPassword : confirmPassword };  // Создаем объект с данными
    const url = 'http://localhost:8080/api/signUp';  // URL-адрес сервера
        
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
                if (response.status == 409){
                    errorMessage.innerHTML = 'This email already exists <br>'
                    return
                }
                errorMessage.innerHTML ='Check everything again <br>';
                return
            } else{
                console.log("okk")
                document.getElementById('authorizationForm').style.display = 'none';
                document.getElementById('succesPage').style.display = 'block';

                
                // Перенаправление на другую страницу
                // window.location.href = "http://localhost:8080/";
            }
        })  // Обрабатываем ответ как JSON

        .catch(error => console.error('Error:', error));  // Обрабатываем возможные ошибки
    
});