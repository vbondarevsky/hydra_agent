;TODO: Это легкий набросок установщика для Windows. Конечно же нужно еще много чего здесь доделать, но это уже что-то :)

[Setup]
AppName=hydra_agent
AppVersion=0.0.1
WizardStyle=modern
DefaultDirName={autopf}\hydra_agent
DefaultGroupName=hydra_agent
UninstallDisplayIcon={app}\hydra_agent.exe
Compression=lzma2
SolidCompression=yes
OutputDir=userdocs:Inno Setup Examples Output
ArchitecturesAllowed=x64
ArchitecturesInstallIn64BitMode=x64

[Files]
Source: "main.exe"; DestDir: "{app}"; DestName: "hydra_agent.exe"
Source: "config.yml"; DestDir: "{app}"; DestName: "config.yml"

[Icons]
Name: "{group}\hydra_agent"; Filename: "{app}\hydra_agent.exe"
