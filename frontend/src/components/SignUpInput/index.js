import React from 'react';

import { Form, Col, Row } from 'react-bootstrap';

export default function SignUpInput({ setState, state, type, placeholder, controlId, label }){

    const handleUser = (e) => {
        console.log(e.target.value);
        setState({...state, [label.toLowerCase()]: e.target.value});
    }

    return (
        <>    
            <Form.Group as={Row} controlId={ controlId }>
                <Form.Label column sm="2">
                    { label }
                </Form.Label>
                <Col sm="10">
                    <Form.Control type={ type } placeholder={ placeholder }  onChange={ handleUser } />
                </Col>
            </Form.Group>
        </>    
    );
}