Pod::Spec.new do |spec|
  spec.name         = 'Gsu'
  spec.version      = '{{.Version}}'
  spec.license      = { :type => 'GNU Lesser General Public License, Version 3.0' }
  spec.homepage     = 'https://github.com/subaru-network/go-subaru'
  spec.authors      = { {{range .Contributors}}
		'{{.Name}}' => '{{.Email}}',{{end}}
	}
  spec.summary      = 'iOS Subaru Client'
  spec.source       = { :git => 'https://github.com/subaru-network/go-subaru.git', :commit => '{{.Commit}}' }

	spec.platform = :ios
  spec.ios.deployment_target  = '9.0'
	spec.ios.vendored_frameworks = 'Frameworks/Gsu.framework'

	spec.prepare_command = <<-CMD
    curl https://gsustore.blob.core.windows.net/builds/{{.Archive}}.tar.gz | tar -xvz
    mkdir Frameworks
    mv {{.Archive}}/Gsu.framework Frameworks
    rm -rf {{.Archive}}
  CMD
end
