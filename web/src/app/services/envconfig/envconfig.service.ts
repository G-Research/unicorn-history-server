/**
 * Licensed to the Apache Software Foundation (ASF) under one
 * or more contributor license agreements.  See the NOTICE file
 * distributed with this work for additional information
 * regarding copyright ownership.  The ASF licenses this file
 * to you under the Apache License, Version 2.0 (the
 * "License"); you may not use this file except in compliance
 * with the License.  You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

import { HttpClient } from '@angular/common/http';
import { Injectable } from '@angular/core';

import { EnvConfig } from '@app/models/envconfig.model';
import { environment } from 'src/environments/environment';

const ENV_CONFIG_JSON_URL = './assets/config/envconfig.json';

export function envConfigFactory(envConfig: EnvConfigService) {
  return () => envConfig.loadEnvConfig();
}

@Injectable({
  providedIn: 'root',
})
export class EnvConfigService {
  private envConfig: EnvConfig;
  private uiProtocol: string;
  private uiHostname: string;
  private uiPort: string;

  constructor(private httpClient: HttpClient) {
    this.uiProtocol = window.location.protocol;
    this.uiHostname = window.location.hostname;
    this.uiPort = window.location.port;
    this.envConfig = {
      yunikornApiURL: 'http://localhost:30001',
      uhsApiURL: 'http://localhost:8989',
    };
  }

  loadEnvConfig(): Promise<void> {
    if (environment.production) {
      console.log('environment.production', environment.production);
      return Promise.resolve();
    }

    return new Promise((resolve) => {
      this.httpClient.get<EnvConfig>(ENV_CONFIG_JSON_URL).subscribe({
        next: (data) => {
          this.envConfig = data;
          resolve();
        },
        error: () => {
          console.warn('Failed to load envconfig.json, using default values');
          resolve();
        },
      });
    });
  }

  getYuniKornWebAddress() {
    if (!environment.production) {
      return `${this.envConfig.yunikornApiURL}/api`;
    }
    return `${this.uiProtocol}//${this.uiHostname}:${this.uiPort}/api`;
  }

  getUHSWebAddress() {
    if (!environment.production) {
      return `${this.envConfig.uhsApiURL}/api`;
    }
    return `${this.uiProtocol}//${this.uiHostname}:${this.uiPort}/api`;
  }

  getExternalLogsBaseUrl() {
    if (!environment.production) {
      return this.envConfig.externalLogsURL || null;
    }
    return null;
  }
}
