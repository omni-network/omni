import React from 'react';
import { useColorMode } from '@docusaurus/theme-common';


function ThemedImage({ lightSrc, darkSrc, alt }) {
    const { colorMode } = useColorMode();
    const src = colorMode === 'dark' ? darkSrc : lightSrc;

    return <img src={src} alt={alt} />;
}

export default ThemedImage;
