import React from 'react';

const Home = (props: { name: string }) => {
    return (
        <div>
            {props.name ? 'Olá ' + props.name : 'Você não está logado'}
        </div>
    );
};

export default Home;