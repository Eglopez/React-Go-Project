import React from 'react';

export default function Input({type, placeholder}){
    return(
        <>
            <input type={type} placeholder={placeholder}/>
        </>
    );
}