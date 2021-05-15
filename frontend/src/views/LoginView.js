import React from 'react';
import Input from '../components/Input';
import LoginForm from '../components/Form';

export default function LoginView(){
    return (
        <>
            <div>
                <LoginForm title="Login">
                    <Input type="text" placeholder="Ingrese @username"/><br></br>
                    <Input type="password" placeholder="Ingrese contraseña"/>
                </LoginForm>
            </div>    
        </>
    );
}