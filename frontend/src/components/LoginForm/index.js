import React from 'react';
import { Button, Form, Container } from 'react-bootstrap';
import './LoginForm.css';
import { Link } from 'react-router-dom';


const LoginForm = ({ children, title, onSubmitFunction }) => {

  
  return (
      <Container className="login-container">
        <Form onSubmit={onSubmitFunction}>
          <h3 className="login-title">{title}</h3>
          { children }
          <Link to="/signup" className="signup-link">Crear Cuenta</Link><br></br>
          <Button type="submit" variant="success" className="btn">Login</Button>
      </Form>
      </Container>
  );
}

export default LoginForm;