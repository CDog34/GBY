import config from 'config';

export class DaoVoiceService {
  static load() {
    if (!config.enableDaoVoice) return null;
    eval(`(function(i,s,o,g,r,a,m){i["DaoVoiceObject"]=r;i[r]=i[r]||function(){(i[r].q=i[r].q||[]).push(arguments)},i[r].l=1*new Date();a=s.createElement(o),m=s.getElementsByTagName(o)[0];a.async=1;a.src=g;a.charset="utf-8";m.parentNode.insertBefore(a,m)})(window,document,"script",('https:' == document.location.protocol ? 'https:' : 'http:') + "//widget.daovoice.io/widget/${config.daoVoiceAppId}.js","daovoice")`);
  }

  static setUpForAnonymous() {
    if (!config.enableDaoVoice) return null;
    window.daovoice('init', {
      app_id: config.daoVoiceAppId
    });
    window.daovoice('update');
  }

  static trackEvent(evtName, evtOpt) {
    if (!config.enableDaoVoice) return null;
    window.daovoice('trackEvent', evtName, {
      appEnv: config.appEnv,
      timeStamp: Date.now(),
      data: evtOpt
    });
  }

  static pageVisitEvent(pageName, opt) {
    const evtOpt = opt || {};
    DaoVoiceService.trackEvent('pageVisit', Object.assign({
      pageName: pageName
    }, evtOpt));
  }

  static outLinkVisitEvent(link) {
    DaoVoiceService.trackEvent('outLinkVisit', link);
  }

  static openNewMsg(msg) {
    if (!config.enableDaoVoice) return null;
    window.daovoice('openNewMessage', msg);
  }
}
