import React from 'react';
import { Button, Form, Col, Row } from 'react-bootstrap';

const LoginForm = ({ children, title, onSubmitFunction }) => {
    return (
        <Form onSubmit={onSubmitFunction}>
            <h3>{title}</h3>
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
    )
}

export default LoginForm;