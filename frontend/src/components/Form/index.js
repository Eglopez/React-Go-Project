import React from 'react';

const Form = ({ children, title, onSubmitFunction }) => {
    return (
        <form onSubmit={onSubmitFunction}>
            <h3>{title}</h3>
            {children}
        </form>
    )
}

export default Form;