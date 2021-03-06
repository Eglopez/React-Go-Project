import React, { useState } from 'react';
import { useHistory } from 'react-router-dom';

import LoginInput from '../components/LoginInput';
import LoginForm from '../components/LoginForm';
import axios from '../axios';

export default function LoginView(){

    const history = useHistory();

    const [user,setUser] = useState({"username":"", "password":""});

    const submitUser = (e) => {
        e.preventDefault();
        console.log(user.password);

        axios({
            url:'http://localhost:4400/api/v1/account/login',
            method:'POST',
            responseType:'json',
            headers:{'Content-Type':'multipart/form-data'},
            data:{
                username: user.username,
                password: user.password
            }
        }).then(res => {
            console.log(res);
            history.push("/signup");
        }).catch(err => {
            console.error(err);
        }); 
        
    } 

    return (
        <>
            <div>
                <LoginForm title="Login" onSubmitFunction={submitUser}>
                   <LoginInput setState={setUser} state={user} type="text" placeholder="@username" controlId="userNameLogin" label="Username" ></LoginInput> 
                   <LoginInput setState={setUser} state={user} type="password" placeholder="Password" controlId="passwordLogin" label="Password" ></LoginInput> 
                </LoginForm>
            </div>    
        </>
    );
}