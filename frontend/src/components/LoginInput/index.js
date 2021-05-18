import React from 'react';
import { Form, Col, Row } from 'react-bootstrap';

export default function Input({ setState, state }){

    const handleUser = (e) => {
        setState({...state, [e.target.name]: e.target.value})
    }

    return(
        <>
            <Form.Group as={Row} controlId="formPlaintextEmail">
                <Form.Label column sm="2">
                    Username
                </Form.Label>
                <Col sm="10">
                    <Form.Control type="text" placeholder="@username" onChange={handleUser}/>
                </Col>
            </Form.Group>
            <Form.Group as={Row} controlId="formPlaintextPassword">
                <Form.Label column sm="2">
                    Password
                </Form.Label>
                <Col sm="10">
                    <Form.Control type="password" placeholder="Password" onChange={handleUser} />
                </Col>
            </Form.Group>
        </>
    );
}