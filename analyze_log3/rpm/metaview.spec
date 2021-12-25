%global metaview_version %{metaview_package_version}


Summary: XTao metaview job scheduler
Name: metaview
Version: %{metaview_version} 
Release: 53.0.0
License: GPL
Buildroot: %{_tmppath}/%{name}-buildroot
Group: Applications/File
%define tarball %{name}-%{version}.tar.gz
Source0: %{tarball}


%description
This is distributed scheduler created by XTao to do large-scale
bioinformatics analysis


%package -n %{name}-server
Summary: metaview scheduler server
Group: System Environment/Kernel
Provides: %{name}-server = %{version}-%{release}

%package -n %{name}-cli
Summary: metaview cli utils
Group: System Environment/Kernel
Provides: %{name}-cli = %{version}-%{release}

%description -n %{name}-server
This is XTao bioinformatic analysis job scheduler

%description -n %{name}-cli
This is XTao bioinformatic analysis CLI utils


%prep
%setup -q


%build
%configure
make

%postun
%define __debug_install_post   \
         %{_rpmconfigdir}/find-debuginfo.sh %{?_find_debuginfo_opts} "%{_builddir}/%{?buildsubdir}"\
         %{nil}

%install
rm -rf %{buildroot}
mkdir -p %{buildroot}/opt/metaview
make DESTDIR=%{buildroot} MANDIR=%{_mandir} BINDIR=%{_sbindir} SYSTEMD_DIR=%{_unitdir} install


%clean
rm -rf %{buildroot}


%files -n %{name}-server
/opt/metaview/metaview

%files -n %{name}-cli
%{_bindir}/*


%changelog
* Thu Feb 16 2017 Jason Zhang <jason.zhang@xtaotech.com> - initial version
- create the initial version

