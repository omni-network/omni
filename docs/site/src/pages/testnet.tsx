// Import React and Layout from Docusaurus
import React from 'react';
import Layout from '@theme/Layout';
import { useColorMode } from '@docusaurus/theme-common';
import CountdownTimer from '../components/CountdownTimer/CountdownTimer';

import './testnet.css';

// Define the ThemedIcon component
const ThemedIcon = () => {
    const { colorMode } = useColorMode();
    const iconSrc = colorMode === 'dark' ? 'img/logo-white.svg' : 'img/logo-dark-blue.svg';

    return <img src={iconSrc} alt="Omni Omega Logo" className="icon" />;
};

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

                {/* Timeline Section */}
                <div className="timeline-section">
                    <h2>Timeline</h2>
                    <p>Omega Testnet will soon be released, stay tuned!</p>
                    <CountdownTimer targetDate="2024-04-05T23:59:59" />
                </div>

                {/* Where to Start Section */}
                <div className="start-section">
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
                            <a href="/operate/testnet/" className="start-box-link"></a>
                            <div className='dot'></div>
                            <ThemedIcon />
                            <h3>Operate</h3>
                            <p>Learn how to start node and join the network as an operator</p>
                        </div>
                    </div>
                </div>

            </div>
        </Layout>
    );
}

// Export the component
export default TestnetPage;
