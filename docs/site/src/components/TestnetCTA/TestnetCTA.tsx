import React from 'react';
import Link from '@docusaurus/Link';

import './testnetcta.css';

const TestnetCTA: React.FC = () => {
  return (
    <Link to="/testnet" className="testnetcta">
      <div className="testnetcta-contents">
        <h2>Join Omni Omega Testnet</h2>
        <p>Be a part of the future of Ethereum interoperability.
          Get started with the Omni Omega Testnet today!</p>
          <div className="animated-text">Get Started with Omega</div>
      </div>
    </Link>
  );
};

export default TestnetCTA;
