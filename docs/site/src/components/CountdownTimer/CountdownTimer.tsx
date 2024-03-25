import { useColorMode } from '@docusaurus/theme-common';
import { CountdownCircleTimer } from "react-countdown-circle-timer";

import "./CountdownTimer.css";

const minuteSeconds = 60;
const hourSeconds = 3600;
const daySeconds = 86400;

const timerProps = {
    isPlaying: true,
    size: 140,
    strokeWidth: 16,
};

const renderTime = (dimension, time) => {
    return (
        <div className="time-wrapper">
            <div className="time"><h3>{time}</h3></div>
            <div className="time-dimension">{dimension}</div>
        </div>
    );
};

const getTimeSeconds = (time) => (minuteSeconds - time) | 0;
const getTimeMinutes = (time) => ((time % hourSeconds) / minuteSeconds) | 0;
const getTimeHours = (time) => ((time % daySeconds) / hourSeconds) | 0;
const getTimeDays = (time) => (time / daySeconds) | 0;

// Accept targetDate as a prop
export default function CountdownTimer({ targetDate }) {
    const endTime = new Date(targetDate).getTime() / 1000; // convert target date to UNIX timestamp in seconds
    const startTime = Date.now() / 1000; // current time in UNIX timestamp seconds
    const remainingTime = endTime - startTime;

    const daysDuration = Math.ceil(remainingTime / daySeconds) * daySeconds;

    const { colorMode } = useColorMode();
    const color = colorMode === 'dark' ? "var(--ifm-heading-color)" : "var(--ifm-heading-color)";
    const colorTrail = colorMode === 'dark' ? "var(--ifm-color-secondary)" : "var(--ifm-color-secondary)";

    return (
        <div className="countdown-timer">
            <CountdownCircleTimer
                {...timerProps}
                colors={color}
                trailColor={colorTrail}
                duration={daysDuration}
                initialRemainingTime={remainingTime}
            >
                {({ elapsedTime }) =>
                    renderTime("Days", getTimeDays(daysDuration - elapsedTime))
                }
            </CountdownCircleTimer>
            <CountdownCircleTimer
                {...timerProps}
                colors={color}
                trailColor={colorTrail}
                duration={daySeconds}
                initialRemainingTime={remainingTime % daySeconds}
                onComplete={(totalElapsedTime) => ({
                    shouldRepeat: remainingTime - totalElapsedTime > hourSeconds,
                })}
            >
                {({ elapsedTime }) =>
                    renderTime("Hours", getTimeHours(daySeconds - elapsedTime))
                }
            </CountdownCircleTimer>
            <CountdownCircleTimer
                {...timerProps}
                colors={color}
                trailColor={colorTrail}
                duration={hourSeconds}
                initialRemainingTime={remainingTime % hourSeconds}
                onComplete={(totalElapsedTime) => ({
                    shouldRepeat: remainingTime - totalElapsedTime > minuteSeconds,
                })}
            >
                {({ elapsedTime }) =>
                    renderTime("Minutes", getTimeMinutes(hourSeconds - elapsedTime))
                }
            </CountdownCircleTimer>
            <CountdownCircleTimer
                {...timerProps}
                colors={color}
                trailColor={colorTrail}
                duration={minuteSeconds}
                initialRemainingTime={remainingTime % minuteSeconds}
                onComplete={(totalElapsedTime) => ({
                    shouldRepeat: remainingTime - totalElapsedTime > 0,
                })}
            >
                {({ elapsedTime }) =>
                    renderTime("Seconds", getTimeSeconds(elapsedTime))
                }
            </CountdownCircleTimer>
        </div>
    );
}
