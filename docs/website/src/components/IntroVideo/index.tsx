import Link from '@docusaurus/Link';
import React, { useEffect, useRef, useState } from 'react';
import styles from './index.module.css';

export default function VideoSequence() {
  const video1Ref = useRef(null);
  const video2Ref = useRef(null);

  useEffect(() => {
    const video1 = video1Ref.current;
    const video2 = video2Ref.current;

    const playVideos = () => {
      video1.play();

      video1.onended = function () {
        video1.style.display = 'none';
        video2.style.display = 'block';
        video2.play();
      };
    };

    // Wait for user interaction before playing the videos
    const handleUserInteraction = () => {
      playVideos();
      document.removeEventListener('click', handleUserInteraction);
      document.removeEventListener('keydown', handleUserInteraction);
    };

    document.addEventListener('click', handleUserInteraction);
    document.addEventListener('keydown', handleUserInteraction);

    return () => {
      document.removeEventListener('click', handleUserInteraction);
      document.removeEventListener('keydown', handleUserInteraction);
    };
  }, []);

  return (
    <div className={styles.videoContainer}>

      <video ref={video1Ref} className={styles.video} controls style={{ display: 'block' }}>
        <source src="/img/omni_1.mp4" type="video/mp4" />
        Your browser does not support the video tag.
      </video>
      <video ref={video2Ref} className={styles.video} loop controls style={{ display: 'none' }}>
        <source src="/img/omni_2.mp4" type="video/mp4" />
        Your browser does not support the video tag.
      </video>
    </div>
  );

}
