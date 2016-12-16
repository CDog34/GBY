import config from 'config';

export class DaoVoiceService {

  static load() {
    eval(`(function(i,s,o,g,r,a,m){i["DaoVoiceObject"]=r;i[r]=i[r]||function(){(i[r].q=i[r].q||[]).push(arguments)},i[r].l=1*new Date();a=s.createElement(o),m=s.getElementsByTagName(o)[0];a.async=1;a.src=g;a.charset="utf-8";m.parentNode.insertBefore(a,m)})(window,document,"script",('https:' == document.location.protocol ? 'https:' : 'http:') + "//widget.daovoice.io/widget/${config.daoVoiceAppId}.js","daovoice")`);
  }

  static setUpForAnonymous() {
    const d = window.daovoice;
    d('init', {
      app_id: config.daoVoiceAppId
    });
    d('update');
  }
}
