import React from 'react';
import { Button, Form, Col, Row, Container } from 'react-bootstrap';

const LoginForm = ({ children, title, onSubmitFunction }) => {
    return (
        <Container className="login-container">
            <Form onSubmit={onSubmitFunction}>
            <h3 className="login-title">{title}</h3>
            <Form.Group as={Row} controlId="formPlaintextEmail">
              <Form.Label column sm="2">
                Username
              </Form.Label>
              <Col sm="10">
                <Form.Control type="text" placeholder="@username" />
              </Col>
            </Form.Group>

            <Form.Group as={Row} controlId="formPlaintextPassword">
              <Form.Label column sm="2">
                Password
              </Form.Label>
              <Col sm="10">
                <Form.Control type="password" placeholder="Password" />
              </Col>
            </Form.Group>
            <Button variant="success" className="btn">Login</Button>
        </Form>
        </Container>
    )
}

export default LoginForm;