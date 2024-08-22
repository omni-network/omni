import clsx from 'clsx';
import Link from '@docusaurus/Link';
import useDocusaurusContext from '@docusaurus/useDocusaurusContext';
import Layout from '@theme/Layout';
import HomepageFeatures from '@site/src/components/HomepageFeatures';
import IntroVideo from '@site/src/components/IntroVideo';
import Heading from '@theme/Heading';

import React, { useEffect, useRef, useState } from 'react';

import styles from './index.module.css';

function HomepageHeader() {
  const {siteConfig} = useDocusaurusContext();

  const video1Ref = useRef(null);
  const video2Ref = useRef(null);
  const [isPlaying, setIsPlaying] = useState(false);

  useEffect(() => {
    const video1 = video1Ref.current;
    const video2 = video2Ref.current;

    if (isPlaying) {
      video1.onended = function () {
        video1.style.display = 'none';
        video2.style.display = 'block';
        video2.play();
      };

      video1.play();
    }
  }, [isPlaying]);

  const handlePlayClick = () => {
    setIsPlaying(true);
  };


  return (
    <header className={clsx('hero', styles.heroBanner)}>
      <div className="container">
        <Heading as="h1" className="hero__title">
        Cross-Chain Liquidity at your fingertips
        </Heading>

       <IntroVideo/>


       <div className={styles.buttons}>
          <Link
            className="button button--secondary button--lg"
            to="/build-xdapp/quickstart">
            Build a dapp in 5min ⏱️
          </Link>
        </div>

      </div>

    </header>
  );
}

export default function Home(): JSX.Element {
  const {siteConfig} = useDocusaurusContext();
  return (
    <Layout
      title={`${siteConfig.title}`}
      description="Omni Developers portal and docs">
      <HomepageHeader />
      <main>
        <HomepageFeatures />
      </main>
    </Layout>
  );
}
