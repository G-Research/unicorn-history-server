import { Component, OnInit } from '@angular/core';

@Component({
  selector: 'app-root',
  templateUrl: './app.component.html',
  styleUrls: ['./app.component.scss'],
})
export class AppComponent implements OnInit {
  title = 'uhs-components';
  remoteEntry: string = '';

  ngOnInit(): void {
    this.remoteEntry = `${window.location.href}remoteEntry.js`;
    console.log(this.remoteEntry);
  }

  copyWebEnvToClipboard() {
    const copyString = `"moduleFederationRemoteEntry": "${this.remoteEntry}",
                "allocationsDrawerRemoteComponent": "./AllocationsDrawerComponent",
                "schedulerService": "./SchedulerService"`;
    navigator.clipboard
      .writeText(copyString)
      .catch((error) => console.error('Writing to the clipboard is not allowed. ', error));
  }
}
