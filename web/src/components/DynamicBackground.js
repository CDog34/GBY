import React from 'react';

import styles from '../styles/components/DynamicBackground.scss'

const BGS = [
  {
    url: 'https://tat.pics/v/1481738565374Konachan.com_-_205563_animal_black_hair_blue_eyes_brown_hair_cape_clouds_fairy_group_leafa_lisbeth_long_hair_male_pina_ponytail_red_hair_sky_sunset_sword_water_weapon_wings.png',
    style: {
      transform: 'translateX(40px) scale(1.1)'
    }
  },
  {
    url: 'https://tat.pics/v/1482512368534555b1aaaa7374.png',
    style: {
      transform: 'translateX(40px) scale(1.1)'
    }
  },
  {
    url: 'https://tat.pics/v/148251411476058427074201509051601333731595451100_027.jpg',
    style: {
      transform: 'translateX(10px) scale(1.05)'
    }
  },
  {
    url: 'https://tat.pics/v/1482514274644bz1080p.com_2016-05-14_18-24-26.jpg',
    style: {
      transform: 'translateX(-40px) scale(1.1)'
    }
  }
];

export class DynamicBackground extends React.Component {
  state = {
    curBg: null,
    bgReady: false
  };

  show() {
    this.setState({bgReady: true});
  }

  async componentWillMount() {
    this.setState({curBg: Math.round(Math.random() * (BGS.length - 1))});
  }

  render() {
    const bg = BGS[this.state.curBg];
    if (!bg) return null;
    return <div className={styles.dynBgWrapper}>
      <img
        src={bg.url}
        onLoad={() => this.show()} className={this.state.bgReady && styles.onShow}
        style={this.state.bgReady ? bg.style : {}}/>
      <div className={styles.bgMask}/>
    </div>
  }
}
