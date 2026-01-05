import { Dispatch, SetStateAction } from "react";

export default class Timer {
    currentTime: number;
    paused: boolean;
    timer: NodeJS.Timeout | null

    constructor() {
        this.currentTime = 0; 
        this.paused = true; 
        this.timer = null;
    }

    start() {
        this.paused = false;
        this.timer = setInterval(() => {
            if(!this.paused) {
                this.currentTime++;
            }
        }, 1000);
    }

    pause() {
        this.paused = true;
        if(this.timer) {
            clearInterval(this.timer);
        }
    }

    reset() {
        this.paused = true; 
        if(this.timer) {
            clearInterval(this.timer);
        }
        this.currentTime = 0; 
    }

    getTime() {
        return this.currentTime;
    }
}