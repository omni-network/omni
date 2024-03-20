// Import React and Layout from Docusaurus
import React from 'react';
import Layout from '@theme/Layout';
import { useColorMode } from '@docusaurus/theme-common';

import './testnet.css';

// Define the ThemedIcon component
const ThemedIcon = () => {
    const { colorMode } = useColorMode();
    const iconSrc = colorMode === 'dark' ? 'img/logo-white.svg' : 'img/logo.svg';

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

                {/* Where to Start Section */}
                <div className="start-section">
                    <h1>Where to Start</h1>
                    <div className="start-container">
                        <div className="start-box">
                            <ThemedIcon />
                            <h3>Use</h3>
                            <p>Learn how to stake <strong>$ETH</strong> and interact with the network</p>
                        </div>
                        <div className="start-box">
                            <ThemedIcon />
                            <h3>Build</h3>
                            <p>Learn how to build native cross-chain applications</p>
                        </div>
                        <div className="start-box">
                            <ThemedIcon />
                            <h3>Operate</h3>
                            <p>Learn how to start node and join the network as an operator</p>
                        </div>
                    </div>
                </div>

                {/* Timeline Section */}
                <div className="timeline-section">
                    <h1>Timeline</h1>
                    <p>Lorem ipsum dolor sit amet, consectetur adipiscing elit.</p>
                </div>

                {/* Reach out to the Community Section */}
                <div className="community-section">
                    <h1>Reach out to the Community</h1>
                    <div className="community-container">
                        <div className="community-box">Twiter</div>
                        <div className="community-box">Discord</div>
                        <div className="community-box">Telegram</div>
                    </div>
                </div>

            </div>
        </Layout>
    );
}

// Export the component
export default TestnetPage;
