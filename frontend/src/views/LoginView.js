import React from 'react';
import Input from '../components/Input';
import Form from '../components/Form';

export default function LoginView(){
    return (
        <>
            <Form title="Login">
                <Input type="text" placeholder="Ingrese @username"/>
                <Input type="password" placeholder="Ingrese contraseña"/>
            </Form>
        </>
    );
}