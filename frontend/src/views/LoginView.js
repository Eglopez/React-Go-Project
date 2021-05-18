import React, { useState } from 'react';
import { useHistory } from 'react-router-dom';

import LoginInput from '../components/LoginInput';
import LoginForm from '../components/LoginForm';

export default function LoginView(){

    const history = useHistory();

    const [user,setUser] = useState({"username":"", "password":""});

    const submitUser = (e) => {
        e.preventDefault();
        console.log(user);
        history.push("/signup");
    } 

    return (
        <>
            <div>
                <LoginForm title="Login" onSubmitFunction={submitUser}>
                   <LoginInput setState={setUser} state={user}></LoginInput> 
                </LoginForm>
            </div>    
        </>
    );
}