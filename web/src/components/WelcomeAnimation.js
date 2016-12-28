import React from 'react';

class animationCore {
  ctx;

  drawingQueue = [];

  constructor(ctx) {
    this.ctx = ctx;
    this.setSize();
    window.addEventListener('resize', () => this.setSize())
  }

  setSize() {
    const cvs = this.ctx.canvas;
    cvs.width = window.innerWidth;
    cvs.height = window.innerHeight;
  }

  drawCircle(props) {
    const {startX, startY, radius, color}=props;
    const ctx = this.ctx;
    ctx.fillStyle = null;
    ctx.beginPath();
    ctx.arc(startX, startY, radius, 0, 2 * Math.PI);
    ctx.fillStyle = color;
    ctx.fill();
  }

  async disappear(duration) {
    duration = duration || 0.5;
    const cvs = this.ctx.canvas;
    cvs.style.transition = `opacity ${duration}s`;
    setTimeout(() => cvs.style.opacity = 0, 100);
    return new Promise((res) => {
      setTimeout(() => {
        res();
      }, duration * 1000 + 100);
    })
  }

  async wait(duration) {
    return new Promise((r) => {
      setTimeout(() => r(), duration);
    })
  }

  async startDrawing() {
    return new Promise((resolve) => {
      const doAnimate = () => {
        let finished = true;
        this.drawingQueue.forEach((drawer) => {
          console.log('[Dbg.jq:drawer.props.color]:', drawer.props.color);
          drawer.render(drawer.props);
          finished = finished && !drawer.update(drawer.props);
        });
        if (finished) return resolve();
        window.requestAnimationFrame(doAnimate.bind(this));
      };
      doAnimate();
    })

  }

  async circleRadiusAnimate(xPos, yPos, startRadius, endRadius, color, speed) {
    return new Promise((resolve) => {
      let currentRadius = startRadius;
      const doAnimate = (radius) => {
        this.drawCircle(xPos, yPos, radius, color);
        currentRadius += speed || 1;
        if (currentRadius > endRadius) return resolve();
        window.requestAnimationFrame(() => doAnimate(currentRadius));
      };
      doAnimate(currentRadius);
    })
  }

  circleRadiusUpdate(end, speed) {
    const spd = speed || 1;
    return (props) => {
      if (props.radius < end) {
        props.radius += spd;
        return true;
      }
      return false;
    }
  }

  async start() {
    // const minLength = window.innerHeight > window.innerWidth ? window.innerWidth : window.innerHeight;
    // await this.circleRadiusAnimate(window.innerWidth / 2, window.innerHeight / 2, 0, minLength * (1 / 6), '#66ccff');
    // await this.wait(1000);
    // await this.disappear(2);
    this.drawingQueue.push({
      render: this.drawCircle.bind(this),
      props: {startX: window.innerWidth / 2, startY: window.innerWidth / 2, radius: 0, color: 'red'},
      update: this.circleRadiusUpdate(400, 1)
    });
    this.drawingQueue.push({
      render: this.drawCircle.bind(this),
      props: {startX: window.innerWidth / 2, startY: window.innerWidth / 2, radius: 0, color: '#66ccff'},
      update: this.circleRadiusUpdate(600, 9)
    });
    await this.startDrawing();
    await this.wait(500);
    await this.disappear(2);
    // this.drawCircle(window.innerWidth / 2, window.innerHeight / 2, minLength * (1 / 3), '#66ccff');
    // return new Promise((resolve) => {
    //   setTimeout(async() => {
    //     await this.disappear(2);
    //     resolve();
    //   }, 5000)
    // })
  }
}

export class WelcomeAnimation extends React.Component {
  canvas;
  animation;
  state = {
    show: true
  };

  async componentDidMount() {
    if (!this.canvas) return null;
    this.animation = new animationCore(this.canvas.getContext('2d'));
    await this.animation.start();
    this.setState({show: false});
  }

  render() {
    if (!this.state.show) return null;
    return <canvas ref={(c) => this.canvas = c}
                   style={{
                     position: 'fixed',
                     top: 0,
                     left: 0,
                     zIndex: 10,
                     backgroundColor: 'white'
                   }}/>
  }
}
