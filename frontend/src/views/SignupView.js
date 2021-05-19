import React, { useState } from 'react';
import { useHistory } from 'react-router-dom';

import SignUpForm from '../components/SignUpForm';
import SignUpInput from '../components/SignUpInput';
import axios from '../axios';

export default function SignupView(){
    
    const history = useHistory();
    const [user,setUser] = useState({"firstname":"", "lastname":"", "email":"", "username":"", "password":""});

    const submitUser = (e) => {
        e.preventDefault();
        console.log(user);

        axios({
            url:'http://localhost:4400/api/v1/account/signup',
            method:'POST',
            responseType:'json',
            headers:{'Content-Type':'multipart/form-data'},
            data:{
               firstname: user.firstname,
               lastname: user.lastname,
               email: user.email,
               username: user.username,
               password: user.password
            }
        }).then(res => {
            console.log(res);
            history.push("/login");
        }).catch(err => {
            console.error(err);
        }); 
    }
  
    
    return (
        <>
            <div>
                <SignUpForm title="Sign Up" onSubmitFunction={submitUser}>
                   <SignUpInput setState={setUser} state={user} type="text" placeholder="Ingrese su nombre" controlId="firstnameSingUP" label="Firstname" ></SignUpInput> 
                   <SignUpInput setState={setUser} state={user} type="text" placeholder="Ingrese su apellido" controlId="lastnameSinUp" label="Lastname" ></SignUpInput>
                   <SignUpInput setState={setUser} state={user} type="email" placeholder="correo electronico" controlId="emailSinUp" label="Email" ></SignUpInput> 
                   <SignUpInput setState={setUser} state={user} type="text" placeholder="Username" controlId="usernameSinUp" label="Username" ></SignUpInput> 
                   <SignUpInput setState={setUser} state={user} type="password" placeholder="Contraseña" controlId="passwordSinUp" label="Password" ></SignUpInput>  
                </SignUpForm>
            </div>    
        </>
    );
}