if [[ "$(whoami)" != "root" ]]; then
        echo "run me with sudo"
        exit 1
fi

echo "Checking for old installs from other sources..."
echo "These can mess up calling go or using features in newer versions..."
if [[ "$(dpkg --list | grep "golang-[0-9].[0-9]" | grep "^ll" | wc -l)" -gt "1" ]]; then
        echo "golang appears to be installed via dpkg."
        read -p "Do you want to uninstall golang via dpkg (Y/N)? " resp
        if [[ "${resp,,}" == "y" ]]; then
                echo "Uninstalling golang via dpkg"
		sudo dpkg -l "golang*" | grep "^ii" | awk '{print $2}' | while read go_install; do 
			dpkg -r "${go_install}"
		done
        else
                echo "Not uninstalling APT's install of golang"
        fi
fi

if [[ "$(apt list --installed | grep "golang-[0-9].[0-9]")" ]]; then
	echo "golang appears to be installed via APT."
	read -p "Do you want to uninstall golang via APT (Y/N)? " resp
	if [[ "${resp,,}" == "y" ]]; then
		echo "Uninstalling golang via APT"
		apt-get remove golang
	else
		echo "Not uninstalling APT's install of golang"
	fi
fi

details=$(curl -s https://go.dev/dl/?mode=json | jq -cr '.[0].files[] | select(.os == "linux") | select(.arch == "amd64")')
filename=$(echo ${details} | jq -r '.filename')
version=$(echo ${details} | jq -r '.version' | sed 's|go||g')
url="https://go.dev/dl/${filename}"

go_install_dir="/usr/local/go"
local_version=$(${go_install_dir}/bin/go version | awk '{print $3}' | sed 's|go||g')
echo "local version:   ${local_version}"
echo "current version: ${version}"
if [[ "${local_version}" == "${version}" ]]; then
	echo "You're already using the latest version."
	echo "Skipping install but checking bash profiles."
else
	# do the installlll
	rm -rf "${go_install_dir}"
	wget ${url}
	tar -C "$(echo ${go_install_dir} | sed 's|/go$||g')" -xzf ${filename}
	chown -R root:root "${go_install_dir}"
	rm ${filename}
fi

[[ "${SUDO_USER}" ]] && user=${SUDO_USER} || user=$(whoami)
bash_path=$(eval echo ~${user})/.bashrc

if [[ ! "$(grep "export GOROOT" ${bash_path})" ]]; then
	echo "export GOROOT=\"${go_install_dir}\"" >> ${bash_path}
	if [[ ! $(grep "export PATH" ${bash_path} | grep "GOROOT") ]]; then
		echo 'PATH=$PATH:$GOROOT/bin' >> ${bash_path}
	fi
fi

if [[ ! $(grep "export GOPATH" ${bash_path}) ]]; then
        echo 'export GOPATH="$HOME/go"' >> ${bash_path}
	if [[ ! $(grep "export PATH" ${bash_path} | grep "GOPATH") ]]; then
                echo 'PATH=$PATH:$GOPATH/bin' >> ${bash_path}
        fi
fi

read -p "Update bash profile for root as well (Y/N) " resp
if [[ "${resp,,}" == "y" ]]; then
	bash_path="/root/.bashrc"
	if [[ ! "$(grep "export GOROOT" ${bash_path})" ]]; then
	        echo "export GOROOT=\"${go_install_dir}\"" >> ${bash_path}
	        if [[ ! $(grep "export PATH" ${bash_path} | grep "GOROOT") ]]; then
	                echo 'PATH=$PATH:$GOROOT/bin' >> ${bash_path}
	        fi
	fi

	if [[ ! $(grep "export GOPATH" ${bash_path}) ]]; then
	        echo 'export GOPATH="$HOME/go"' >> ${bash_path}
	        if [[ ! $(grep "export PATH" ${bash_path} | grep "GOPATH") ]]; then
	                echo 'PATH=$PATH:$GOPATH/bin' >> ${bash_path}
	        fi
	fi
fi

echo "Finsihed install. Don't forget to do \"source ~/.bashrc\" to reload bash profile"
