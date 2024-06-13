// Кидаем запрос на /api/profile и исходя от ответа отоброжаем логин или логаут
// Тут происходит создание кнопок login , logout
async function createLoginLogout() {
    let statusLogin = false;
    try {
        const response = await fetch("http://localhost:8080/api/profile", {
            method: 'GET',
            headers: {
                'Content-Type': 'application/json'
            },
        });

        if (!response.ok) {
            if (response.status == 401) {
                statusLogin = false;
            } else {
                throw new Error('Network response was not ok ' + response.statusText);
            }
        } else {
            // Обработка успешного ответа, если нужно
            statusLogin = true
            const data = await response.json();
            // Работа с данными профиля
        }
    } catch (error) {
        console.error('Ошибка при обновлении реакции:', error);
    }

    const result = document.createElement('li');
    if (statusLogin == true){
        result.textContent = "Logout";
        result.addEventListener('click',() => {
            logoutFetch()
        });
    } else{
        result.textContent = "Login";
        result.addEventListener('click',() => window.open("http://localhost:8080/signIn"));
    }
    document.getElementById('headList').appendChild(result);
}
function logoutFetch(){
    // Отправляем запрос на сервер
        fetch("http://localhost:8080/api/logout", {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json'
            },
        })
        .then(response => {
            if (!response.ok){
                throw new Error('Network response was not ok ' + response.statusText);
            } else{
                // Перезагружаем страницу
                window.location.href = "http://localhost:8080/";
            }
        })  // Обрабатываем ответ как JSON

        .catch(error => console.error('Error:', error));  // Обрабатываем возможные ошибки
}

// Скрывает и показывает див при нажатии
function toggleDiv(id){
    const div = document.getElementById(id);
    div.style.display = div.style.display == "flex" ? "none" : "flex";
}
// При выботе фильтра
function handleSelectChange() {
    const selectElement = document.getElementById('selectFilterPost');
    const selectedValue = selectElement.value;
    console.log('Selected value:', selectedValue);
    let url = `http://localhost:8080/api/posts/all`
    // код для обработки выбора
    if (selectedValue === "liked" || selectedValue === "disliked"){
        url = `http://localhost:8080/api/posts/all?filter=${selectedValue}`
    } else if (selectedValue === "Art" || selectedValue === "Technology" || selectedValue === "Science"){
        url = `http://localhost:8080/api/posts/filter/category?name=${selectedValue}`
    }
    document.getElementById("postContainer").innerHTML = ""
    fetchJson(url)
    .then(result => {
       if(result.status != 200){
            throw new Error('Network response was not ok ' + response.statusText);
        }
        result.data.forEach(post => {
            createPost("postContainer",post.createdBy,post.id, post.userId, post.createdAt, post.title, post.content, post.category, post.likes, post.dislikes, post.liked, post.disliked);
        });
    })
    .catch(error => {
        console.error('Ошибка при получении данных:', error);
    });


}

// Функция для создания  постов из ответа сервера
function createPost(divId,username, id, userId, createdTime, title, content, category, likeCount, dislikeCount, liked, disliked) {
    const postDiv = document.createElement('div');
    postDiv.classList.add('post', "centerBlock");
    postDiv.id = `post_${id}`;

    const postHead = document.createElement('div');
    postHead.className = 'postHead';
    postHead.innerHTML = `<b onclick= "createdByPosts(${userId})" >${username} </b><br> ${createdTime}`;
    postDiv.appendChild(postHead);

    const postTitle = document.createElement('h3');
    postTitle.textContent = title;
    postDiv.appendChild(postTitle);

    const postContent = document.createElement('h4');
    postContent.textContent = content;
    postDiv.appendChild(postContent);

    const postCategory = document.createElement('i');
    postCategory.textContent = "Category: " + category;
    postDiv.appendChild(postCategory);

    const postBottom = createPostBottom(likeCount, dislikeCount, liked,disliked)
    postDiv.appendChild(postBottom);


    const commentArea = document.createElement('div');
    commentArea.className = 'commentArea';
    postDiv.appendChild(commentArea);

    const inputMsg = document.createElement('div');
    inputMsg.className = 'inputMsg';
    inputMsg.innerHTML = `
    <textarea class="comment_input" id="comment_input_${id}" placeholder="Введите комментарий"></textarea>
    <button class="submit_comment" onclick="sendComment(${id})">Отправить</button>
    `
    postDiv.appendChild(inputMsg);
    document.getElementById(divId).appendChild(postDiv);
}


//создание нижней части поста (лайк дизлайк сообщение)
// А сами кнопки создаются в другой функции  createReactionButton
function createPostBottom(likeCount, dislikeCount, liked, disliked){
    const postBottom = document.createElement('div');
    postBottom.className = 'postBottom';

    const likeButton = createReactionButton('like', likeCount, liked);
    postBottom.appendChild(likeButton);

    const dislikeButton = createReactionButton('dislike', dislikeCount, disliked);
    postBottom.appendChild(dislikeButton);

    const commentsButton = createReactionButton('comments', "...", false);
    commentsButton.classList.add("commentBtn");
    postBottom.appendChild(commentsButton);

    return postBottom
}

// Тут создаются кнопки с функиями ОНКЛИК
function createReactionButton(type, count, status) {
    const buttonDiv = document.createElement('div');
    buttonDiv.className = 'reactionBtns';
    if (status == true){
        buttonDiv.classList.add('liked');
    }
    buttonDiv.dataset.type = type;
    buttonDiv.dataset.active = 'false';

    let svgIcon;
    if (type === 'like') {
        svgIcon = '<svg width="30px" height="30px" viewBox="0 0 24 24" fill="none" xmlns="http://www.w3.org/2000/svg" transform="matrix(1, 0, 0, -1, 0, 0)"><path d="M20.2699 8.48505L20.9754 12.5648C21.1516 13.5838 20.368 14.5158 19.335 14.5158H14.1539C13.6404 14.5158 13.2494 14.9767 13.3325 15.484L13.9952 19.5286C14.1028 20.1857 14.0721 20.858 13.9049 21.5025C13.7664 22.0364 13.3545 22.465 12.8128 22.6391L12.6678 22.6856C12.3404 22.7908 11.9831 22.7663 11.6744 22.6176C11.3347 22.4539 11.0861 22.1553 10.994 21.8001L10.5183 19.9663C10.3669 19.3828 10.1465 18.8195 9.86218 18.2874C9.44683 17.5098 8.80465 16.8875 8.13711 16.3123L6.69838 15.0725C6.29272 14.7229 6.07968 14.1994 6.12584 13.6656L6.93801 4.27293C7.0125 3.41139 7.7328 2.75 8.59658 2.75H13.2452C16.7265 2.75 19.6975 5.17561 20.2699 8.48505Z" fill="currentColor"/><path fill-rule="evenodd" clip-rule="evenodd" d="M2.96767 15.2651C3.36893 15.2824 3.71261 14.9806 3.74721 14.5804L4.71881 3.34389C4.78122 2.6221 4.21268 2 3.48671 2C2.80289 2 2.25 2.55474 2.25 3.23726V14.5158C2.25 14.9174 2.5664 15.2478 2.96767 15.2651Z" fill="currentColor"/></svg>';
    } else if (type === 'dislike') {
        svgIcon = '<svg width="30px" height="30px" viewBox="0 0 24 24" fill="none" xmlns="http://www.w3.org/2000/svg"><path d="M20.2699 8.48505L20.9754 12.5648C21.1516 13.5838 20.368 14.5158 19.335 14.5158H14.1539C13.6404 14.5158 13.2494 14.9767 13.3325 15.484L13.9952 19.5286C14.1028 20.1857 14.0721 20.858 13.9049 21.5025C13.7664 22.0364 13.3545 22.465 12.8128 22.6391L12.6678 22.6856C12.3404 22.7908 11.9831 22.7663 11.6744 22.6176C11.3347 22.4539 11.0861 22.1553 10.994 21.8001L10.5183 19.9663C10.3669 19.3828 10.1465 18.8195 9.86218 18.2874C9.44683 17.5098 8.80465 16.8875 8.13711 16.3123L6.69838 15.0725C6.29272 14.7229 6.07968 14.1994 6.12584 13.6656L6.93801 4.27293C7.0125 3.41139 7.7328 2.75 8.59658 2.75H13.2452C16.7265 2.75 19.6975 5.17561 20.2699 8.48505Z" fill="currentColor"/><path fill-rule="evenodd" clip-rule="evenodd" d="M2.96767 15.2651C3.36893 15.2824 3.71261 14.9806 3.74721 14.5804L4.71881 3.34389C4.78122 2.6221 4.21268 2 3.48671 2C2.80289 2 2.25 2.55474 2.25 3.23726V14.5158C2.25 14.9174 2.5664 15.2478 2.96767 15.2651Z" fill="currentColor"/></svg>';
    } else if (type === 'comments') {
        svgIcon = '<svg width="40px" height="40px" viewBox="0 -0.5 25 25" fill="white" xmlns="http://www.w3.org/2000/svg"><path fill-rule="evenodd" clip-rule="evenodd" d="M18.1 5.00016H6.9C6.53425 4.99455 6.18126 5.13448 5.9187 5.38917C5.65614 5.64385 5.50553 5.99242 5.5 6.35816V14.5002C5.50553 14.8659 5.65614 15.2145 5.9187 15.4692C6.18126 15.7238 6.53425 15.8638 6.9 15.8582H10.77C10.9881 15.857 11.2035 15.9056 11.4 16.0002L17.051 19.0002L17 14.5002H18.43C19.0106 14.5091 19.4891 14.0467 19.5 13.4662V6.35816C19.4945 5.99242 19.3439 5.64385 19.0813 5.38917C18.8187 5.13448 18.4657 4.99455 18.1 5.00016Z" stroke="currentColor" stroke-width="1.5" stroke-linecap="round" stroke-linejoin="round"/><path d="M8.5 8.25024C8.08579 8.25024 7.75 8.58603 7.75 9.00024C7.75 9.41446 8.08579 9.75024 8.5 9.75024V8.25024ZM16.5 9.75024C16.9142 9.75024 17.25 9.41446 17.25 9.00024C17.25 8.58603 16.9142 8.25024 16.5 8.25024V9.75024ZM8.5 11.2502C8.08579 11.2502 7.75 11.586 7.75 12.0002C7.75 12.4145 8.08579 12.7502 8.5 12.7502V11.2502ZM14.5 12.7502C14.9142 12.7502 15.25 12.4145 15.25 12.0002C15.25 11.586 14.9142 11.2502 14.5 11.2502V12.7502ZM8.5 9.75024H16.5V8.25024H8.5V9.75024ZM8.5 12.7502H14.5V11.2502H8.5V12.7502Z" fill="currentColor"/></svg>';
    }

    buttonDiv.innerHTML = `${svgIcon} <span>${count}</span>`;
    buttonDiv.addEventListener('click', handleReactionClick);
    return buttonDiv;
}

// Функция обновляет количество лайков и дизлайков
async function updatePostBottom(id) {
    let postBottom;

    try {
        const response = await fetch(`http://localhost:8080/api/posts/one?valueId=${id}`, {

        });
        if (!response.ok) {
            throw new Error('Network response was not ok ' + response.statusText);
        }

        const data = await response.json();

        postBottom = createPostBottom(data.likes, data.dislikes, data.liked, data.disliked);
    } catch (error) {
        console.error('There was a problem with the fetch operation:', error);
        return;
    }

    const post = document.getElementById(`post_${id}`);
    // Найти внутри элемента 'post' элемент с классом 'bottom'
    const bottom = post.querySelector('.postBottom');
    // Заменяем второй див на новый див
    post.replaceChild(postBottom, bottom);
}


//Реагирует на нажатие сабмит
document.getElementById('newPostForm').addEventListener('submit', function(event) {
    event.preventDefault();
    const title = document.getElementById('TitlePost').value;
    const content = document.getElementById('contentPost').value;
    const category = document.getElementById('selectCategoryPost').value;

    const data = { Title : title, Content : content, Category : category};  // Создаем объект с данными
    const url = 'http://localhost:8080/api/posts';  // URL-адрес сервера
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
                throw new Error('Network response was not ok ' + response.statusText);
            } else{
                // Перенаправление на другую страницу
                window.location.href = "http://localhost:8080/";
            }
        })  // Обрабатываем ответ как JSON

        .catch(error => console.error('Error:', error));  // Обрабатываем возможные ошибки

});

// Обработчик событий для кнопок реакций
function handleReactionClick(event) {
    const button = event.currentTarget;
    const postDiv = button.closest('.post'); // Найти ближайший элемент с классом 'post'
    const postId = postDiv.id; // Получить id поста

    const reactionType = button.dataset.type; // Тип реакции (лайк, дизлайк или комментарии)
    let arr = postId.split('_');

    let url
    if (reactionType == 'like'){
        url = `http://localhost:8080/api/posts/like?valueId=${arr[1]}`
    } else if (reactionType == 'dislike'){
        url = `http://localhost:8080/api/posts/dislike?valueId=${arr[1]}`
    } else if(reactionType == "commentLike"){
        const commentDiv = button.closest('.commentBlock'); // Найти ближайший элемент с классом 'post'
        const commentId = commentDiv.id; // Получить id поста
        let comment_id = commentId.split('_');
        url = `http://localhost:8080/api/comments/like?valueId=${comment_id[1]}`
    } else if(reactionType == "commentDislike"){
        const commentDiv = button.closest('.commentBlock'); // Найти ближайший элемент с классом 'post'
        const commentId = commentDiv.id; // Получить id поста
        let comment_id = commentId.split('_');
        url = `http://localhost:8080/api/comments/dislike?valueId=${comment_id[1]}`
    } else {
        postDiv.querySelector('.inputMsg').style.display = "block";
        createCommentArea(arr[1])
        return
    }
    console.log(url)
    // Отправляем запрос на сервер
    fetch(url, {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json'
        },
    })
    .then(response => {
        if (!response.ok) {
            if (response.status == 401){
                alert("register to leave reactions")
                return
            }
            throw new Error('Network response was not ok ' + response.statusText);
        } else{
            updatePostBottom(arr[1])
            if (reactionType == "commentLike" || reactionType == "commentDislike"){
                createCommentArea(arr[1])
            }
        }
        return response.json();
    })
    .catch(error => console.error('Ошибка при обновлении реакции:', error));
}


// Место которое будет вставлено коментарии  с сервера
function createCommentArea(id){
    // Создаем Див для коммента и даю класс
    const commentArea = document.createElement('div');
    commentArea.className = 'commentArea';

    fetch(`http://localhost:8080/api/comments?valueId=${id}`)
    .then(response => {
        if (!response.ok) {
            console.log("no comments:", response.status)
        }
        return response.json();
    })
    .then(data => {
        if(data.message == "Nothing found... :|"){
            console.log("no comments:")
            return
        }
        data.forEach(post => {
            const commentBlock = document.createElement('div');
            commentBlock.className = 'commentBlock';
            commentBlock.id = `comment_${post.id}`
            commentBlock.innerHTML = `
                <b>${post.createdBy}</b> </br> <span>${post.content}</span> </br>
            `

            commentBlock.appendChild(commentReaction(post.liked, post.likes, "commentLike"))
            commentBlock.appendChild(commentReaction(post.disliked, post.dislikes, "commentDislike"))
            commentArea.appendChild(commentBlock);
    });
  })
  .catch(error => {
        console.log("comment load", error)
        return
  });
    // Добавляю готовую в Пост
    // document.getElementById(`post_${id}`).appendChild(commentArea);
    const post = document.getElementById(`post_${id}`);
    // Найти внутри элемента 'post' элемент с классом 'bottom'
    const old = post.querySelector('.commentArea');
    // Заменяем второй див на новый див
    commentArea.style.display = "block";
    post.replaceChild(commentArea, old);
}

// Отправка комента на сервер
function sendComment(id){
    const commentInput = document.getElementById(`comment_input_${id}`)
    let commentMsg = commentInput.value
    // Проверяем, что комментарий не пустой
    if (commentMsg.trim() !== '') {
        const data = {Content : commentMsg};  // Создаем объект с данными
        const url = `http://localhost:8080/api/comments/post?valueId=${id}`  // URL-адрес сервера
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
                throw new Error('Network response was not ok ' + response.statusText);
            } else{
                createCommentArea(id)
            }
        })  // Обрабатываем ответ как JSON
        .catch(error => console.error('Error:', error));  // Обрабатываем возможные ошибки
        commentInput.value = ''; // Очищаем поле ввода после отправки
    } else {
        console.log('Пустой комментарий не может быть отправлен');
    }
}

function commentReaction(status, count, type){
    const reactionBtn = document.createElement('b');
    if (type == "commentLike"){
        reactionBtn.textContent = `Liked: ${count}`;
    } else if (type == "commentDislike"){
        reactionBtn.textContent = `Disiked: ${count}`;
    } else{
        console.log("Неправильный вызов функции commentReaction")
        return
    }
    reactionBtn.dataset.type = type;
    if (status == true){
        reactionBtn.className = "liked"
    }
    reactionBtn.addEventListener('click', handleReactionClick);
    return reactionBtn
}

function profileInfo(id, login, email){
    const profleInfoBlock = document.getElementById("profileInfo");
    const userLogin = document.createElement('b');
    userLogin.textContent = `Login: ${login}`;
    profleInfoBlock.appendChild(userLogin);
    const userEmail = document.createElement('b');
    userEmail.textContent = `Email: ${email}`;
    profleInfoBlock.appendChild(userEmail);
}

function createdByPosts(id){
    document.getElementById("addPost").style.display = "none";
    document.getElementById("filterArea").style.display = "none";

    document.getElementById("postContainer").innerHTML= '';
    fetch(`http://localhost:8080/api/posts/createdBy?valueId=${id}`)
    .then(response => {
    if (!response.ok) {
      throw new Error('Network response was not ok ' + response.statusText);
    } else{
    }
    return response.json();
  })
  .then(data => {
    data.forEach(post => {
        createPost("postContainer",post.createdBy,post.id, post.userId, post.createdAt, post.title, post.content, post.category, post.likes, post.dislikes, post.liked, post.disliked);

    });
  })
  .catch(error => {
    console.error('There was a problem with the fetch operation:', error);
  });
}


function myLikedPosts(id) {
    // отправляем запрос на мною созданных постов
    fetchJson(`http://localhost:8080/api/posts/createdBy?valueId=${id}`)
    .then(result => {
        if(result.status != 200){
            console.log('Network response was not ok ' + result.status);
            return
        }
        console.log('Статус ответа:', result.status);
        result.data.forEach(post => {
            createPost("myPosts",post.createdBy,post.id, post.userId, post.createdAt, post.title, post.content, post.category, post.likes, post.dislikes, post.liked, post.disliked);
        });
    })
    .catch(error => {
        console.error('Ошибка при получении данных:', error);
    });
}

// При загрузке страницы
window.addEventListener('load', () => {
    createLoginLogout()
    const savedSection = localStorage.getItem('selectedSection') || 'home';
    if (savedSection == "profile"){
        ProfilePage()
    } else if (savedSection == "contacts"){
        contactsPage()
    } else{
        homePage()
    }

});

// При загрузке гланого экрана
function homePage(){
    document.getElementById("homeBtn").style.fontWeight = "bold";
    document.getElementById("profilePage").style.display = "none"
    document.getElementById("contactsPage").style.display = "none";

    fetch('http://localhost:8080/api/posts/all')
    .then(response => {
        if (!response.ok) {
        throw new Error('Network response was not ok ' + response.statusText);
        }
        return response.json();
    })
    .then(data => {
        data.forEach(post => {
            createPost("postContainer",post.createdBy,post.id, post.userId, post.createdAt, post.title, post.content, post.category, post.likes, post.dislikes, post.liked, post.disliked);
        });
    })
    .catch(error => {
        console.error('There was a problem with the fetch operation:', error);
    });
}

// При загрузке Profile
function ProfilePage(){
    document.getElementById("ProfileBtn").style.fontWeight = "bold";
    document.getElementById("homePage").style.display = "none";
    document.getElementById("contactsPage").style.display = "none";

    // отправляем запрос на получение данных пользователя
    fetchJson('http://localhost:8080/api/profile')
    .then(result => {
        if (result.status === 401){
            localStorage.clear();
            // отправляем на регистрацию
            window.location.href = "http://localhost:8080/signUp";
            return
        } else if(result.status != 200){
            throw new Error('Network response was not ok ' + response.statusText);
        }
        myLikedPosts(result.data.id)
        profileInfo(result.data.id, result.data.login, result.data.email);
    })
    .catch(error => {
        console.error('Ошибка при получении данных:', error);
    });

    // отправляем запрос на получение лайкнутых постов
    fetchJson('http://localhost:8080/api/posts/liked')
    .then(result => {
        console.log('Статус ответа:', result.status);
        result.data.forEach(post => {
            createPost("likedPosts",post.createdBy,post.id, post.userId, post.createdAt, post.title, post.content, post.category, post.likes, post.dislikes, post.liked, post.disliked);
        });
    })
    .catch(error => {
        console.error('Ошибка при получении данных:', error);
    });
}
// При загрузке Contacts
function contactsPage(){
    document.getElementById("homePage").style.display = "none";
    document.getElementById("profilePage").style.display = "none";
    document.getElementById("contactsBtn").style.fontWeight = "bold";
}
// Функция которая отправляет ГЕТ запрос на указанный url
async function fetchJson(url) {
    try {
        const response = await fetch(url);
        const status = response.status;
        const data = await response.json();
        return { data, status };
    } catch (error) {
        throw error;
    }
}

function changePage(event){
    const sectionId = event.target.getAttribute('data-section');
    localStorage.setItem('selectedSection', sectionId);
    // Перезагружаем страницу
    window.location.href = "http://localhost:8080/";
}

function Burger(){
    document.getElementById("headList").classList.toggle('active');
}