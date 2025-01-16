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

const fs = require('fs');

let config = process.argv[2];
let template_environment = fs.readFileSync('./src/environments/environment.template.ts').toString();
let envconfig = JSON.parse(fs.readFileSync('./src/assets/config/envconfig.json').toString());

Object.keys(process.env).forEach((env_var) => {
  // Handle TypeScript environment file
  let regex = new RegExp(`(${env_var}:)(\\s*?)('.*?'|".*?")`);
  template_environment = template_environment.replace(
    regex,
    `${env_var}: '${process.env[env_var]}'`
  );

  // Handle JSON config file
  if (envconfig.hasOwnProperty(env_var)) {
    envconfig[env_var] = process.env[env_var];
  }
});

if (config && config === 'prod')
  fs.writeFileSync(`./src/environments/environment.${config}.ts`, template_environment);
else fs.writeFileSync('./src/environments/environment.ts', template_environment);

fs.writeFileSync('./src/assets/config/envconfig.json', JSON.stringify(envconfig, null, 2));
