// Import React and Layout from Docusaurus
import React from 'react';
import Layout from '@theme/Layout';
import { useColorMode } from '@docusaurus/theme-common';

import './testnet.css';

// Define the ThemedIcon component
const ThemedIcon = () => {
    const { colorMode } = useColorMode();
    const iconSrc = colorMode === 'dark' ? 'img/logo-white.svg' : 'img/logo-dark-blue.svg';

    return <img src={iconSrc} alt="Omni Omega Logo" className="icon" />;
};

const XLogo = () => (
    <svg viewBox="0 0 15 14" fill="none" xmlns="http://www.w3.org/2000/svg">
        <path d="M.534 0 5.94 7.227.5 13.103h1.224L6.486 7.96l3.848 5.144H14.5L8.79 5.47 13.855 0H12.63L8.244 4.738 4.7 0H.534Zm1.8.902h1.914l8.452 11.3h-1.914L2.334.902Z"></path>
    </svg>
);

const DiscordLogo = () => (
    <svg viewBox="0 0 16 16" fill="none" xmlns="http://www.w3.org/2000/svg">
        <path d="M13.545 2.66a12.745 12.745 0 0 0-3.257-1.06.049.049 0 0 0-.052.026 9.84 9.84 0 0 0-.406.875 11.613 11.613 0 0 0-3.658 0 9.032 9.032 0 0 0-.412-.875.05.05 0 0 0-.052-.026 12.709 12.709 0 0 0-3.257 1.06.047.047 0 0 0-.021.02C.356 5.932-.213 9.105.066 12.238c.001.016.01.03.02.04a13.038 13.038 0 0 0 3.996 2.12.05.05 0 0 0 .056-.02c.308-.441.582-.906.818-1.395a.054.054 0 0 0-.028-.074 8.564 8.564 0 0 1-1.248-.625.055.055 0 0 1-.005-.089c.084-.066.168-.135.248-.204a.048.048 0 0 1 .051-.007c2.619 1.255 5.454 1.255 8.041 0a.048.048 0 0 1 .053.006c.08.07.164.139.248.205a.055.055 0 0 1-.004.09 8.038 8.038 0 0 1-1.249.623.055.055 0 0 0-.027.075c.24.488.514.953.817 1.394a.05.05 0 0 0 .056.02 12.994 12.994 0 0 0 4.001-2.12.055.055 0 0 0 .021-.038c.334-3.622-.559-6.769-2.365-9.558a.042.042 0 0 0-.021-.02Zm-8.198 7.67c-.789 0-1.438-.76-1.438-1.692 0-.933.637-1.693 1.438-1.693.807 0 1.45.767 1.438 1.693 0 .933-.637 1.693-1.438 1.693Zm5.316 0c-.788 0-1.438-.76-1.438-1.692 0-.933.637-1.693 1.438-1.693.807 0 1.45.767 1.438 1.693 0 .933-.63 1.693-1.438 1.693Z">
        </path>
    </svg>
);

const TelegramLogo = () => (
    <svg viewBox="0 0 16 16" fill="none" xmlns="http://www.w3.org/2000/svg">
        <path d="M1.1 7.11c4.295-1.806 7.159-2.997 8.592-3.573 4.091-1.642 4.941-1.928 5.496-1.937.121-.002.394.027.57.165.15.117.19.275.21.385.02.11.044.363.025.56-.222 2.249-1.181 7.706-1.67 10.225-.206 1.066-.613 1.423-1.006 1.458-.856.076-1.505-.546-2.334-1.07-1.297-.82-2.03-1.331-3.288-2.132-1.455-.925-.512-1.434.317-2.265.217-.218 3.987-3.527 4.06-3.828.01-.037.018-.177-.069-.251-.086-.074-.213-.049-.305-.029-.13.029-2.201 1.35-6.214 3.965-.588.39-1.12.58-1.598.57-.526-.01-1.538-.287-2.29-.523C.673 8.54-.06 8.387.004 7.896.037 7.64.402 7.378 1.1 7.11Z" clip-rule="evenodd">
        </path>
    </svg>
);

// Define the TestnetPage component
function TestnetPage() {
    return (
        <Layout
            title="Omni Omega Testnet Participation"
            description="Learn how to participate in the Omni Omega testnet phases.">
            <div className="container">
                <br />

                {/* Banner Section */}
                <div className="banner">
                    <h1 className="banner-title">Omni Omega</h1>
                    <h2>Welcome to the Omni Omega Testnet Home!</h2>
                    <h4 className="banner-subtitle">Find here all there is to know about being part of this phase of the network.</h4>
                </div>

                {/* Where to Start Section */}
                <div className="start-section">
                    <h1>Where to Start</h1>
                    <div className="start-container">
                        <div className="start-box">
                            <a href="/learn/testnet/" className="start-box-link"></a>
                            <div className='dot'></div>
                            <ThemedIcon />
                            <h3>Use</h3>
                            <p>Learn how to stake <strong>$ETH</strong> and interact with the network</p>
                        </div>
                        <div className="start-box">
                            <a href="/develop/testnet/" className="start-box-link"></a>
                            <div className='dot'></div>
                            <ThemedIcon />
                            <h3>Build</h3>
                            <p>Learn how to build native cross-chain applications</p>
                        </div>
                        <div className="start-box">
                            <a href="/operate/onboarding/" className="start-box-link"></a>
                            <div className='dot'></div>
                            <ThemedIcon />
                            <h3>Operate</h3>
                            <p>Learn how to start node and join the network as an operator</p>
                        </div>
                    </div>
                </div>


                {/* Timeline Section */}
                <div className="timeline-section">
                    <h1>Timeline</h1>
                    <p>⏳ A timeline for Omega Testnet will soon be released, stay tuned! ⏳</p>
                </div>

                {/* Reach out to the Community Section */}
                <div className="community-section">
                    <h1>Reach out to the Community</h1>
                    <div className="community-container">
                        <a href="https://twitter.com/OmniFDN" className="community-box"><XLogo /></a>
                        <a href="https://discord.gg/bKNXmaX9VD" className="community-box"><DiscordLogo /></a>
                        <a href="https://t.me/omnifdn" className="community-box"><TelegramLogo /></a>
                    </div>
                </div>

            </div>
        </Layout>
    );
}

// Export the component
export default TestnetPage;
