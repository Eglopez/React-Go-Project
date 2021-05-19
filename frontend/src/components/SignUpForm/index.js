import React, { useState } from 'react';
import { Button, Form, Container } from 'react-bootstrap';
import './SignUpForm.css';
import { Link } from 'react-router-dom';

const SignUpForm = ({ children, title, onSubmitFunction }) => {
    return (
        <Container className="signUp-container">
            <Form onSubmit={onSubmitFunction}>
                <h3 className="SignUp-title">{title}</h3>
                { children }
                <Link to="/login" className="login-link">Iniciar Sesion</Link><br></br>
                <Button type="submit" variant="success" className="btn">SingUp</Button>
            </Form>
        </Container>
    );
}

export default SignUpForm;